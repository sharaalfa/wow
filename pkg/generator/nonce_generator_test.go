package generator

import (
	"testing"
	"time"
)

func TestCheckHashPrefix(t *testing.T) {
	tests := []struct {
		name       string
		hash       []byte
		difficulty int
		expected   bool
	}{
		{
			name:       "All zero bytes, difficulty 8",
			hash:       []byte{0x00, 0x00, 0x00},
			difficulty: 8,
			expected:   true,
		},
		{
			name:       "All zero bytes, difficulty 16",
			hash:       []byte{0x00, 0x00, 0x00},
			difficulty: 16,
			expected:   true,
		},
		{
			name:       "Non-zero first byte, difficulty 8",
			hash:       []byte{0x01, 0x00, 0x00},
			difficulty: 8,
			expected:   false,
		},
		{
			name:       "Non-zero second byte, difficulty 16",
			hash:       []byte{0x00, 0x01, 0x00},
			difficulty: 16,
			expected:   false,
		},
		{
			name:       "Zero bits in last byte, difficulty 17",
			hash:       []byte{0x00, 0x00, 0x00},
			difficulty: 17,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckHashPrefix(tt.hash, tt.difficulty)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestGenerateValidNonce(t *testing.T) {
	tests := []struct {
		name      string
		challenge string
		expected  string
	}{
		{
			name:      "Valid challenge",
			challenge: "123456:3",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout := time.After(5 * time.Second)
			done := make(chan string)

			go func() {
				nonce := GenerateValidNonce(tt.challenge)
				done <- nonce
			}()

			select {
			case nonce := <-done:
				if tt.expected == "" && nonce == "" {
					t.Errorf("Expected non-empty nonce, got empty")
				} else if tt.expected != "" && nonce != tt.expected {
					t.Errorf("Expected nonce %s, got %s", tt.expected, nonce)
				}
			case <-timeout:
				t.Fatal("Test timed out. Could not generate a valid nonce within the time limit.")
			}
		})
	}
}
