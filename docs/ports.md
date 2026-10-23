# Ports

Claimed from the third-party-network range (`22600–22699`, one block of ten per network), after
Among Us (22600–22609), `servers/demonware` (22610–22619), `servers/eos` (22620–22629),
`servers/t17` (22630–22639), `platforms/switch/games/klei` (22640–22649), `servers/playfab` (22650–22659)
and `platforms/switch/games/spelunky2` (22660–22669).

| Port | Use |
| --- | --- |
| 22670 | HTTP for the Shiro backend names (behind Traefik, which terminates the guest's TLS) |
| 22671 | `/health` |
| 22672–22679 | reserved for this backend (lobby/ranked/relay, once the capture shows what it needs) |

One port serves every API path until the bytecode says otherwise: paths are
distinguished by path, not by port, and Traefik routes each Shiro name to this
service. Relay/transport ports will claim reserved ports here once measured —
see `protocol.md` §4.
