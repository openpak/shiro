# shiro

OpenPak's backend for Shiro Games' Switch titles: the custom HTTPS service
Northgard dials once the console's DNS points at OpenPak. One repository for
the vendor, like `platforms/switch/games/klei` and `servers/t17`, with titles as records
under `docs/`.

## Status: dev-deployed (2026-10-13) — `shirod` serves and logs, Northgard first

`openpak-shiro` runs on the box (`localhost/openpak-shiro:dev`, internal
22670 + health 22671, verified: `/health` ok, unknown paths 404 + redacted
log). No Traefik router and no ceiling claim yet — both are owner steps (CA
key for `shiro.pem`, ceiling key for v8); the CSR and the runbook are staged
(`next-session.md`).

| Title | Status |
| --- | --- |
| Northgard (0100A9E00D97A000, v1.2.0) | stack proven custom (no NEX/NPLN/LDN — `docs/protocol.md`); Shiro hostnames still in the romfs bytecode (`docs/evidence-needed.md`) |

`cmd/shirod`, one Go binary, plain HTTP on 22670 (health 22671) behind a
Traefik TCP router that terminates the console's TLS for the Shiro names.
Until the bytecode yields the real paths, every request is answered 404 and
logged redacted — the log is the spec for the first handlers.

```sh
go run ./cmd/shirod            # :22670, health :22671
```

## Ground rules

- Clean room. Public protocol documentation, our own captures and
  observations, and compatibly licensed code are the only inputs.
- AGPL-3.0-only, English only, Go 1.26, no secrets or production identifiers
  in the repository. Logs carry shapes, never tokens, UIDs, names or bodies
  (`redact.go`).

This project is not affiliated with or endorsed by Shiro Games or Nintendo.
It contains no game code, assets, SDK material, binaries, certificates,
keys, credentials or raw captures.
