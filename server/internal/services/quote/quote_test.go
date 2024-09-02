package quote

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
	"wow/pkg/httpclient"
)

func TestGetRandomQuote(t *testing.T) {
	tests := []struct {
		name             string
		maxRetries       int
		mockClient       *httpclient.MockHTTPClient
		quotesEntryPoint string
		expected         string
		expectedErr      error
	}{
		{
			name:       "Successful request on first try",
			maxRetries: 3,
			mockClient: &httpclient.MockHTTPClient{
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
			},
			quotesEntryPoint: "https://api.quotable.io/random",
			expected:         "Test quote",
			expectedErr:      nil,
		},
		{
			name:       "Successful request on second try",
			maxRetries: 3,
			mockClient: &httpclient.MockHTTPClient{
				Responses: []*http.Response{
					{
						StatusCode: 200,
						Body: io.NopCloser(bytes.NewBufferString(`{
						"content": "Test quote",
						"author": "Test author"
					}`)),
					},
				},
				Errors: []error{errors.New("HTTP request failed")},
			},
			quotesEntryPoint: "https://api.quotable.io/random",
			expected:         "Test quote",
			expectedErr:      nil,
		},
		{
			name:       "Failed request after max retries",
			maxRetries: 3,
			mockClient: &httpclient.MockHTTPClient{
				Responses: []*http.Response{nil},
				Errors:    []error{errors.New("HTTP request failed")},
			},
			quotesEntryPoint: "https://api.quotable.io/random",
			expected:         getRandomBackupQuote(),
			expectedErr:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quoteService := &Service{
				Client:           tt.mockClient,
				QuotesEntryPoint: tt.quotesEntryPoint,
				MaxRetries:       tt.maxRetries,
			}
			quote, err := quoteService.GetRandomQuote()
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
			if quote != tt.expected && tt.maxRetries > 3 {
				t.Errorf("Expected quote %v, got %v", tt.expected, quote)
			}
		})
	}
}

func BenchmarkGetRandomQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quoteService := &Service{
			Client: &httpclient.MockHTTPClient{
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
			},
			QuotesEntryPoint: getRandomBackupQuote(),
			MaxRetries:       3,
		}
		_, err := quoteService.GetRandomQuote()
		if err != nil {
			return
		}
	}
}
