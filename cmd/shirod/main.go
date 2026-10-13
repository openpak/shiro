// Command shirod serves Shiro Games' own online backend for their Switch
// titles, Northgard (0100A9E00D97A000, current update v1.2.0) first.
//
// Northgard on Switch is neither NEX nor NPLN: its native main links only
// nn::socket and nn::ssl for transport, nn::account for the Nintendo Account
// ID token, and nn::friends for profile lookup. There is no NEX game server
// ID, no NPLN tenant, no LDN, no Photon, no EOS and no PlayFab marker in the
// base build — see docs/protocol.md. The game logic runs on the HashLink VM
// (northgard.nss bytecode), so the Shiro hostnames and API shapes live in the
// romfs bytecode, not in the native main, and are still unknown.
//
// The one route every Shiro title is known to need is the Nintendo token
// check: the game reads the nn::account token cache
// (EnsureNetworkServiceAccountAvailable,
// LoadNetworkServiceAccountIdTokenCache) and the real backend has Nintendo
// vouch for it. Ours will ask nx-baas's /internal/switch/identity instead —
// an OpenPak token is ours to accept — once the capture shows which field
// carries it.
//
// Plain HTTP inside the pod network; Traefik terminates the guest's TLS for
// the Shiro names and forwards the stream, the same arrangement as
// servers/klei, servers/t17, servers/spelunky2 and the Among Us front.
//
// Every endpoint below is a logging stub until the first Harbor boot teaches
// us the request shapes; an unhandled request is logged in full (redacted)
// because the request is the specification for its handler.
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	addr       = flag.String("addr", ":22670", "HTTP listen address (Traefik terminates TLS in front)")
	healthAddr = flag.String("health-addr", ":22671", "health listen address")
)

func main() {
	flag.Parse()
	s := newServer(newIdentityClient())

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.route)
	go func() {
		h := http.NewServeMux()
		h.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
			writeJSON(w, 200, map[string]any{"status": "ok"})
		})
		log.Printf("[shiro] health on %s", *healthAddr)
		log.Fatal(http.ListenAndServe(*healthAddr, h))
	}()

	log.Printf("[shiro] serving Shiro backend on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}

type server struct {
	ids identityClient
}

func newServer(ids identityClient) *server { return &server{ids: ids} }

// route dispatches on path first and hostname second. Until the romfs
// bytecode yields the real Shiro hostnames and paths, everything lands in
// the default branch and is logged for the protocol spec.
func (s *server) route(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if i := strings.IndexByte(host, ':'); i >= 0 {
		host = host[:i]
	}
	switch {
	default:
		body, _ := io.ReadAll(io.LimitReader(r.Body, 1<<16))
		log.Printf("[shiro] !! unhandled %s %s%s body=%s", r.Method, host, logURL(r.URL), redactBody(body))
		writeJSON(w, 404, map[string]any{"error": "not found"})
	}
}
