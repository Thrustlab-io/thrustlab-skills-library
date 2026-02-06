package clay

import (
	"crypto/rand"
	"fmt"
)

func GenerateSessionID() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		randomBytes(4), randomBytes(2), randomBytes(2), randomBytes(2), randomBytes(6))
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
