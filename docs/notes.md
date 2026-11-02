# Notes: shiro

## Overview

OpenPak's backend for Shiro Games' Switch titles: the custom HTTPS service
Northgard dials once the console's DNS points at OpenPak. One repository for
the vendor, like `platforms/switch/games/klei` and `servers/t17`, with titles as records
under `docs/`.

## Service

`openpak-shiro` (`shirod`): internal port 22670, health 22671. Unknown paths answer 404 and log redacted.
The `shiro.pem` leaf comes from the OpenPak CA and the hostnames go through the signed ceiling.

| Title | Notes |
| --- | --- |
| Northgard (0100A9E00D97A000, v1.2.0) | custom stack (no NEX/NPLN/LDN — `docs/protocol.md`); Shiro hostnames are in the romfs bytecode (`docs/evidence-needed.md`) |

## Ground rules

- Clean room. Public protocol documentation, our own captures and
  observations, and compatibly licensed code are the only inputs.
- AGPL-3.0-only, English only, Go 1.26, no secrets or production identifiers
  in the repository. Logs carry shapes, never tokens, UIDs, names or bodies
  (`redact.go`).

This project is not affiliated with or endorsed by Shiro Games or Nintendo.
It contains no game code, assets, SDK material, binaries, certificates,
keys, credentials or raw captures.
