package handler

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
	g "wow/pkg/generator"
	"wow/pkg/handler"
	"wow/pkg/httpclient"
	"wow/server/internal/services/generator"
	"wow/server/internal/services/quote"
)

func TestHandleConnection(t *testing.T) {
	challenge := "123456:3"
	validNonce := g.GenerateValidNonce(challenge)

	tests := []struct {
		name           string
		difficulty     int
		timeout        time.Duration
		readData       string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Successful connection",
			difficulty:     3,
			timeout:        10 * time.Second,
			readData:       validNonce,
			expectedOutput: "Test quote",
			expectedError:  nil,
		},
		{
			name:           "Invalid nonce",
			difficulty:     3,
			timeout:        10 * time.Second,
			readData:       "invalid_nonce",
			expectedOutput: "Invalid solution",
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockConn := &handler.MockConn{
				ReadBuffer:  bytes.NewBufferString(tt.readData),
				WriteBuffer: &bytes.Buffer{},
			}

			quoteService := quote.NewService(&httpclient.MockHTTPClient{
				Responses: []*http.Response{
					{
						StatusCode: 200,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"content": "Test quote",
							"author": "Test author"
						}`)),
					},
				},
				Errors: []error{nil},
			}, "https://api.quotable.io/random", 3)

			ch := NewConnectionHandler(tt.difficulty, tt.timeout, quoteService)
			mockHandleConnection(ch, mockConn, validNonce)

			output := mockConn.WriteBuffer.String()
			lines := strings.Split(output, "\n")
			lastLine := lines[len(lines)-2]

			if lastLine != tt.expectedOutput {
				t.Errorf("Expected output %q, got %q", tt.expectedOutput, lastLine)
			}
		})
	}
}

func mockHandleConnection(ch *ConnectionHandler, conn net.Conn, validNonce string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close connection: ", err.Error())
		}
	}(conn)

	challenge := generator.GenerateChallenge(ch.Difficulty)
	if err := ch.sendChallenge(conn, challenge); err != nil {
		log.Println("Error sending challenge: ", err.Error())
		return
	}

	nonce, err := ch.readOnce(conn)
	if err != nil {
		log.Println("Error reading nonce: ", err.Error())
		return
	}

	if nonce == validNonce {
		if err := ch.sendQuote(conn); err != nil {
			log.Println("Error sending quote: ", err.Error())
		}
	} else {
		if err := ch.sendInvalidSolution(conn); err != nil {
			log.Println("Error sending invalid solution: ", err.Error())
		}
	}
}
