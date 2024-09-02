package httpclient

import "net/http"

type MockHTTPClient struct {
	Responses []*http.Response
	Errors    []error
	callCount int
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	if m.callCount >= len(m.Responses) {
		m.callCount = 0
	}
	resp := m.Responses[m.callCount]
	err := m.Errors[m.callCount]
	m.callCount++
	return resp, err

}
