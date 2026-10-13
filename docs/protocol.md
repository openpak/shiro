# Shiro backend protocol (Northgard Switch, 0100A9E00D97A000)

Sources are our own observations of a legitimately owned copy (base + update
v1.2.0); no vendor code, no captures from other projects. Clean-room rules
apply: public docs, our own captures, compatibly licensed code only.

## 1. Stack verdict: custom Shiro backend, not NEX/NPLN

Base `main.bin` (26MB, extracted from the owned base NSP) carries none of the
known stack markers (`catalogue/netscan.py`: NEX, Pia, NPLN, Photon, EOS,
Demonware, Battle.net, Unity-Services, PlayFab, Nakama, Steam, LDN) and no
hostname strings at all.

Its `nn::` module surface is the whole story:

| Module | Symbols | Meaning |
| --- | --- | --- |
| `socket` | 17 | raw TCP (`Connect`, `Send`, `Recv`, `Select`, …) |
| `ssl` | 16 | TLS (`Connection::DoHandshake`, `Read`, `Write`, `SetHostName`) |
| `account` | 17 | user select, NSO license check, NSA ID token cache |
| `friends` | 5 | `GetProfileList` by NSA ID (friend display, not matchmaking) |
| `nifm` | 6 | network availability |
| `ec` | 1 | `ShowShopAddOnContentList` (DLC shop) |

Absent: `nex`, `pia`, `npln`, `ldn`, `http`, `bcat`, `notification`. There is
no NEX game server ID and no NPLN tenant for this title.

## 2. Auth shape (from imports, pending capture confirmation)

* `EnsureNetworkServiceAccountAvailable` + `ShowLicenseRequirementsForNetworkService`
  — linked Nintendo Account and NSO membership required (matches the eShop
  listing: online 1–6, NSO required).
* `LoadNetworkServiceAccountIdTokenCache` /
  `EnsureNetworkServiceAccountIdTokenCacheAsync` — the game reads the NSA ID
  token and hands it to Shiro's backend, which has Nintendo vouch for it.
  Ours will verify via nx-baas `/internal/switch/identity` (EOS precedent).

## 3. Engine: HashLink VM

The native main is a runner; game logic is HashLink bytecode:

* `D:/Projects/shiroTools/hashlink/…` build paths in the binary
* `D:\Projects\northgard_console\bin\nx\NX64\Release\northgard.nss`

Consequence: Shiro hostnames and API shapes live in the romfs bytecode, not
in the native main. The native main is host-free by construction.

## 4. Hostnames (CONFIRMED 2026-10-13, Harbor)

* `master.shirogames.com` — the game's only server contact, once at boot:
  `sfdnsres` resolve, one stream socket, one `nn::ssl` connection. No new
  traffic when the multiplayer menu opens; the menu reuses the boot session.
* **Port 8080, not 443** (probed 2026-10-13: every other port refused on
  178.33.233.205, 8080 open): TLS 1.3, no ALPN, HTTP/1.1. `GET /` answers
  `401 Unauthorized`, `WWW-Authenticate: Basic realm="Please identify
  yourself"` — the API sits behind HTTP Basic auth. Harbor's earlier
  `send errno 7` failures were connections to local stand-ins with nothing
  on 8080, not a game or emulator bug.
* Unclaimed at observation time, so it resolved to the real address
  (178.33.233.205) and TLS went to Shiro, not to us.
* Claim path: `openpak-policy ceiling add switch .shirogames.com`, then the
  platforms.yml entry, the Traefik SNI route to `shirod:22670`, and the
  testdata envelope refresh in the same change.

## 5. Unknowns (see evidence-needed.md)

* Shiro hostnames and API paths (romfs bytecode extraction pending).
  `master.shirogames.com` confirmed 2026-10-13 (§4); relay hosts, if any,
  appear only in real gameplay and are still unknown.
* Whether update v1.2.0 adds any `nn::` module — CLOSED 2026-10-13: the
  update's exefs main is byte-identical to base (`md5 79624ac42b8df953986ece576214e42f`,
  26,164,240 bytes, no markers). v1.2.0 changes romfs only; the §1 census
  covers the running build completely.
* Rendezvous vs. relay shape: the eShop page claims "dedicated servers";
  ranked GP ladder endpoint unknown.

## 6. Gameplay relay: via servers/connectivity, not custom ports

Shiro claims no relay ports in 22672–22679. When captures show the gameplay
transport needing NAT traversal, `shirod` registers each player as a
connectivity participant (namespace `shiro`, title `northgard`, per-connection
generation), serves TURN credentials from `/v1/ice` mapped into Shiro's own
session-join responses, and allocates UDP pair relays via `/v1/relays` once
the game's native handshake is understood well enough to verify endpoints.
The typed client lives in `servers/connectivity/client/`; the backend token
is an operator-provisioned secret (box connectivity config), never in this
repo. Until the transport is measured, all of 22672–22679 stays reserved and
unspent.
