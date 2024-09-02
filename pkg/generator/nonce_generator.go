package generator

import (
	"crypto/sha256"
	"strconv"
	"strings"
)

func CheckHashPrefix(hash []byte, difficulty int) bool {
	for i := 0; i < difficulty/8; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	remainder := difficulty % 8
	if remainder != 0 {
		mask := byte(0xFF) >> remainder
		if hash[difficulty/8]&mask != 0 {
			return false
		}
	}

	return true
}

func GenerateValidNonce(challenge string) string {
	parts := strings.Split(challenge, ":")
	if len(parts) != 2 {
		return ""
	}

	data := parts[0]
	difficulty, _ := strconv.Atoi(parts[1])

	for i := 0; ; i++ {
		nonce := strconv.Itoa(i)
		hash := sha256.Sum256([]byte(data + nonce))
		if CheckHashPrefix(hash[:], difficulty) {
			return nonce
		}
	}
}
