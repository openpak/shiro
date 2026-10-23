# syntax=docker/dockerfile:1
# Shiro Games' own online backend (Northgard on Switch first): lobby/ranked
# rendezvous plus the Nintendo token check. Plain HTTP inside the pod
# network — Traefik terminates the guest's TLS for the Shiro names, the same
# arrangement as platforms/switch/games/klei, servers/t17, platforms/switch/games/spelunky2 and the Among
# Us front.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -buildvcs=false -ldflags="-s -w" -o /out/shirod ./cmd/shirod

FROM gcr.io/distroless/static
COPY --from=build /out/shirod /shirod
USER nonroot:nonroot
EXPOSE 22670 22671
ENTRYPOINT ["/shirod"]
