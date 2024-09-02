package verifier

import (
	"testing"
	"time"
	"wow/pkg/generator"
)

func TestVerifySolution(t *testing.T) {
	challenge := "123456:3"

	timeout := time.After(5 * time.Second)
	done := make(chan string)

	go func() {
		validNonce := generator.GenerateValidNonce(challenge)
		done <- validNonce
	}()

	var validNonce string
	select {
	case validNonce = <-done:
	case <-timeout:
		t.Fatal("Test timed out. Could not generate a valid nonce within the time limit.")
	}

	tests := []struct {
		name      string
		challenge string
		nonce     string
		expected  bool
	}{
		{
			name:      "Valid nonce",
			challenge: challenge,
			nonce:     validNonce,
			expected:  true,
		},
		{
			name:      "Invalid nonce",
			challenge: challenge,
			nonce:     "invalid_nonce",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := VerifySolution(tt.challenge, tt.nonce)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func BenchmarkVerifySolution(b *testing.B) {
	challenge := "1234567890:4"
	nonce := "12345"
	for i := 0; i < b.N; i++ {
		VerifySolution(challenge, nonce)
	}
}
