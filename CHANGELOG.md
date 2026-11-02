# Changelog

## v0.1.0 — 2026-11-02

- Release workflow publishing to ghcr.io; replaces the hand-built production image.

- Add a release workflow publishing the image to ghcr.io on v*.*.* tags.


- Scaffold `shirod`: plain-HTTP logging server on 22670, `/health` on 22671,
  redacting log helpers, stub nx-baas identity client.
- `docs/protocol.md`: Northgard Switch stack verdict (custom Shiro backend
  over `nn::socket`/`ssl` + NSA ID token; no NEX/NPLN/LDN), HashLink bytecode
  finding, open unknowns.
- Port block 22670–22679 claimed in `docs/ports.md` and the global port plan.
