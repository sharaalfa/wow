package quote

import (
	"math/rand"
	"time"
	"wow/pkg/httpclient"
)

type Service struct {
	Client           httpclient.HTTPClient
	QuotesEntryPoint string
	MaxRetries       int
}

func NewService(client httpclient.HTTPClient, quotesEntrypoint string, maxRetries int) *Service {
	return &Service{
		Client:           client,
		QuotesEntryPoint: quotesEntrypoint,
		MaxRetries:       maxRetries,
	}

}

func (s *Service) GetRandomQuote() (string, error) {
	for i := 0; i < s.MaxRetries; i++ {
		quote, err := RequestRandomQuote(s.Client, s.QuotesEntryPoint)
		if err == nil {
			return quote.Content, nil
		}
	}
	return getRandomBackupQuote(), nil
}

func getRandomBackupQuote() string {
	quotes := []string{
		"The only way to do great work is to love what you do. - Steve Jobs",
		"Believe you can and you're halfway there. - Theodore Roosevelt",
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return quotes[rand.Intn(len(quotes))]
}
