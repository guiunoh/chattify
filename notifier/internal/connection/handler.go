package connection

import (
	"context"
	"log"
	"notifier/internal/connection/adaptor"
	"notifier/pkg/ulid"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Command struct {
	Type  string `json:"type"`
	Topic string `json:"topic"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type Handler struct {
	hub       adaptor.Hub
	presenter Presenter
	usecase   Usecase
	commands  map[string]func(c context.Context, connector adaptor.Connector, topic string)
}

func NewHandler(hub adaptor.Hub, usecase Usecase) *Handler {
	handler := Handler{
		usecase:   usecase,
		hub:       hub,
		commands:  make(map[string]func(c context.Context, connector adaptor.Connector, topic string)),
		presenter: NewPresenter(),
	}

	return &handler
}

func (h *Handler) Route(r gin.IRoutes) {
	r.GET("/ws", h.handle)
}

func (h *Handler) handle(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	cCp := c.Copy()
	id := h.usecase.GenerateID().String()
	connector := adaptor.NewConnector(id, conn)

	connector.SetOnConnectHandler(func(connector adaptor.Connector) {
		h.hub.Register(connector)
		h.OnConnect(cCp, connector)
	})

	connector.SetOnCloseHandler(func(connector adaptor.Connector) {
		h.hub.UnRegister(connector)
		h.OnClose(cCp, connector)
	})

	connector.SetOnReceiveHandler(func(connector adaptor.Connector, message []byte) {
		h.OnReceive(cCp, connector, message)
	})

	connector.SetOnPingHandler(func(connector adaptor.Connector) {
		h.OnPing(cCp, connector)
	})

	connector.Accept()
}

func (h *Handler) OnConnect(c *gin.Context, connector adaptor.Connector) {
	log.Println("OnConnect:", connector.ID(), c.RemoteIP())
	id, err := ulid.ParseID(connector.ID())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	conn, err := h.usecase.CreateConnection(ctx, id)
	if err != nil {
		panic(err)
	}

	connector.SendMessage(h.presenter.Connected(conn.ID, conn.ExpiryAt))
}

func (h *Handler) OnClose(c *gin.Context, connector adaptor.Connector) {
	log.Println("OnClose:", connector.ID(), c.RemoteIP())

	id, err := ulid.ParseID(connector.ID())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err = h.usecase.CloseConnection(ctx, id); err != nil {
		panic(err)
	}
}

func (h *Handler) OnPing(c *gin.Context, connector adaptor.Connector) {
	log.Println("onPing")
	ctx := context.Background()

	id, err := ulid.ParseID(connector.ID())
	if err != nil {
		panic(err)
	}

	if err := h.usecase.ExtendConnection(ctx, id); err != nil {
		return
	}
}

func (h *Handler) OnReceive(c *gin.Context, connector adaptor.Connector, message []byte) {
	//handler.commands["forwarder"] = handler.onSubscribe
	//handler.commands["unsubscribe"] = handler.onUnSubscribe

	//var cmd Command
	//if err := json.Unmarshal(message, &cmd); err != nil {
	//	connector.SendMessage(h.presenter.Usage())
	//	panic(err)
	//}
	//
	//fn, ok := h.commands[cmd.Type]
	//if !ok {
	//	connector.SendMessage(h.presenter.Usage())
	//	panic(fmt.Sprintf("unknown command:%s", cmd.Type))
	//}
	//fn(context.Background(), connector, cmd.Topic)
	connector.SendMessage(string(message))
}
