package quote

import (
	"encoding/json"
	"io"
	"log"
	"wow/pkg/httpclient"
)

type Response struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Quote struct {
	Content string
	Author  string
}

func RequestRandomQuote(client httpclient.HTTPClient, quotesEntryPoint string) (Quote, error) {
	resp, err := client.Get(quotesEntryPoint)
	if err != nil {
		return Quote{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close body: ", err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quote{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Quote{}, err
	}
	return Quote{
		Content: response.Content,
		Author:  response.Author,
	}, nil
}
