package quote

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
	"wow/pkg/httpclient"
)

func TestRequestRandomQuote(t *testing.T) {
	tests := []struct {
		name             string
		mockClient       *httpclient.MockHTTPClient
		quotesEntryPoint string
		expected         Quote
		expectedErr      error
	}{
		{
			name: "Successful request",
			mockClient: &httpclient.MockHTTPClient{
				Responses: []*http.Response{
					{
						StatusCode: 200,
						Body: io.NopCloser(bytes.NewBufferString(
							`{
								 "content": "Test quote",
								 "author": "Test author"
								}`)),
					},
				},
				Errors: []error{nil},
			},
			quotesEntryPoint: "https://api.quotable.io/random",
			expected: Quote{
				Content: "Test quote",
				Author:  "Test author",
			},
			expectedErr: nil,
		},
		{
			name: "Failed request",
			mockClient: &httpclient.MockHTTPClient{
				Responses: []*http.Response{nil},
				Errors:    []error{errors.New("HTTP request failed")},
			},
			quotesEntryPoint: "https://api.quotable.io/random",
			expected:         Quote{},
			expectedErr:      errors.New("HTTP request failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quote, err := RequestRandomQuote(tt.mockClient, tt.quotesEntryPoint)
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
			if quote != tt.expected {
				t.Errorf("Expected quote %v, got %v", tt.expected, quote)
			}
		})
	}
}
