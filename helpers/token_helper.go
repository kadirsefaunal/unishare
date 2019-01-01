package helpers

import (
	"crypto/sha1"
	"encoding/hex"
)

func CreateTokenByUserName(userName string) string {
	// Todo: Change the token creation way.
	h := sha1.New()
	h.Write([]byte(userName))
	token := hex.EncodeToString(h.Sum(nil))

	return token
}
