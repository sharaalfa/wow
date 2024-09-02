package verifier

import (
	"crypto/sha256"
	"strconv"
	"strings"
	"wow/pkg/generator"
)

func VerifySolution(challenge, nonce string) bool {
	parts := strings.SplitN(challenge, ":", 2)
	if len(parts) != 2 {
		return false
	}

	data := parts[0] + nonce
	hash := sha256.Sum256([]byte(data))

	difficulty, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}

	return generator.CheckHashPrefix(hash[:], difficulty)

}
