package forwarder

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func NewHandler(usecase Usecase, rdb *redis.Client, httpclient *http.Client, channel string) Handler {
	return &handler{usecase, rdb, httpclient, channel}
}

type handler struct {
	usecase    Usecase
	rdb        *redis.Client
	httpclient *http.Client
	channel    string
}

func (h *handler) RunLoop() {
	ctx := context.Background()
	pubsub := h.rdb.Subscribe(ctx, h.channel)
	defer pubsub.Close()

	channel := pubsub.Channel()
	for message := range channel {
		// HTTP 요청 생성
		go func(data string) {
			defer func() {
				if r := recover(); r != nil {
					log.Println("receive panic:", r)
				}
			}()
			h.subscribe([]byte(data))
		}(message.Payload)
	}
}

func (h *handler) subscribe(data []byte) {
	var input struct {
		ConnectionID string `json:"connectionID"`
		Payload      string `json:"payload"`
	}

	if err := json.Unmarshal(data, &input); err != nil {
		log.Println("error parsing json:", err)
		return
	}

	if err := h.usecase.ReceiveMessage(input.ConnectionID, input.Payload); err != nil {
		log.Println("error usecase receive message:", err)
	}
}
