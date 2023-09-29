package utils

import (
	"crypto/sha256"
	"encoding/hex"
	// ... other imports ...
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
