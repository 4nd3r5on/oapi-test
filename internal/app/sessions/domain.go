package sessions

import (
	"crypto/rand"
	"encoding/hex"
)

func NewKey() (string, error) {
	b := make([]byte, KeyLen)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
