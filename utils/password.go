package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GeneratePassword(rut string) string {
	newPass := strings.ReplaceAll(strings.ReplaceAll(rut, ".", ""), "-", "")
	newPassBytes := sha256.Sum256([]byte(newPass))
	return hex.EncodeToString(newPassBytes[:])
}
