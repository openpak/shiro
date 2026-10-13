# Evidence needed

1. **Shiro hostnames + API shapes.** Extract the romfs from the owned update
   NSP (`Northgard [0100A9E00D97A800][v524288][Up v1.2.0].nsp`), find the
   HashLink bytecode (`northgard.nss` or equivalent blob), and read out the
   hostname strings and path/payload shapes. Native `main.bin` is proven
   host-free, so the bytecode is the only source.
2. **Update module census.** Decompress the update's exefs main and confirm
   the `nn::` module list matches base (socket/ssl/account/friends, no
   nex/npln/ldn). Extraction was blocked by host IO pressure on 2026-10-13
   (nstool stuck in disk-wait); retry with `catalogue/exefs.sh`.
3. **First Harbor boot.** Boot Northgard on Harbor with DNS pointed at
   OpenPak, open the multiplayer menu, and record the DNS names + the
   redacted request log from `shirod`. The request log is the spec for the
   first real handlers (lobby list, token check, ranked).
4. **Ranked/GP endpoint.** The eShop page advertises ranks; find which call
   carries the GP update and whether it shares the lobby host.
