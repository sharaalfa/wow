package generator

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateChallenge(t *testing.T) {
	difficulty := 4
	challenge := GenerateChallenge(difficulty)
	parts := strings.Split(challenge, ":")
	if len(parts) != 2 {
		t.Errorf("Challenge format is incorrect: %s", challenge)
	}
	_, err := strconv.Atoi(parts[1])
	if err != nil {
		t.Errorf("Difficulty is not a number: %s", parts[1])
	}
}

func BenchmarkGenerateChallenge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateChallenge(4)
	}
}
