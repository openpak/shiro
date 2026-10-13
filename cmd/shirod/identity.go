package main

// identityClient verifies a Nintendo Account ID token against nx-baas. The
// stub below answers "unknown" for everything until the first capture shows
// which field carries the token; then Verify will call
// nx-baas's /internal/switch/identity, the same way servers/eos verifies the
// NSA id_token for Connect login.
type identityClient struct {
	nxInternalURL string
	nxInternalKey string
}

func newIdentityClient() identityClient {
	return identityClient{}
}

// Verify reports whether token belongs to a real OpenPak-linked user. Stub:
// always unknown until wired to nx-baas.
func (c identityClient) Verify(token string) (uid string, ok bool) {
	_ = token
	return "", false
}
