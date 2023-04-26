package forwarder

import (
	"bytes"
	"log"
	"mime"
	"net/http"

	"github.com/goccy/go-json"
)

func NewGateway(client *http.Client, endpoint string) Gateway {
	return &gateway{client, endpoint}
}

type gateway struct {
	client   *http.Client
	endpoint string
}

func (g gateway) PostTo(connectionID string, data string) error {
	params := struct {
		ConnectionID string `json:"connectionID"`
		Payload      string `json:"payload"`
	}{connectionID, data}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	resp, err := g.client.Post(g.endpoint, mime.TypeByExtension(".json"), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("HTTP response status code:", resp.StatusCode)
	return nil
}
