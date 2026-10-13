package main

import (
	"net/url"
	"strings"
)

// logPath cuts a path to a sane length for the log. No secrets are expected
// in Shiro paths, but length is bounded anyway.
func logPath(p string) string {
	if len(p) > 256 {
		return p[:256] + "…"
	}
	return p
}

// redactQuery renders query values with anything token-shaped hidden. Names
// are kept so the spec can record which field carries what.
func redactQuery(vals url.Values) string {
	var b strings.Builder
	first := true
	for k, vs := range vals {
		for _, v := range vs {
			if !first {
				b.WriteByte('&')
			}
			first = false
			b.WriteString(k)
			b.WriteByte('=')
			if looksSecret(k, v) {
				b.WriteString("<redacted>")
			} else {
				b.WriteString(v)
			}
		}
	}
	return b.String()
}

// redactBody hides bodies that are opaque blobs or carry tokens, and cuts the
// rest. JSON bodies are logged in full only when short; the shapes go into
// docs/protocol.md, never raw tokens.
func redactBody(body []byte) string {
	s := string(body)
	if len(s) > 512 {
		s = s[:512] + "…"
	}
	for _, kw := range []string{"token", "ticket", "auth", "password", "secret"} {
		if strings.Contains(strings.ToLower(s), kw) {
			return "<withheld: carries credentials>"
		}
	}
	return s
}

func looksSecret(k, v string) bool {
	kl := strings.ToLower(k)
	for _, kw := range []string{"token", "ticket", "auth", "secret", "password", "code"} {
		if strings.Contains(kl, kw) {
			return true
		}
	}
	return len(v) > 64
}
