# Deploy: shiro

`cmd/shirod`, one Go binary, plain HTTP on 22670 (health 22671) behind a
Traefik TCP router that terminates the console's TLS for the Shiro names.
Until the bytecode yields the real paths, every request is answered 404 and
logged redacted — the log is the spec for the first handlers.

```sh
go run ./cmd/shirod            # :22670, health :22671
```

## Image

Built from the repo `Dockerfile`: a `golang:1.26-alpine` build stage produces
a static `/shirod` binary, which runs on `gcr.io/distroless/static` as
`nonroot:nonroot`. It exposes 22670 and 22671 and uses `/shirod` as the
entrypoint.

## Ports

See [ports.md](ports.md).
