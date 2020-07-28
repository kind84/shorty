package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	encoded := base64.RawURLEncoding.EncodeToString([]byte(sum))
	return encoded[:10]
}
