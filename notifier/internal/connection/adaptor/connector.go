package adaptor

import (
	"log"
	"time"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

const (
	liveCheckPeriod = 10 * time.Minute
	pingPeriod      = 60 * time.Second
	writeWait       = 3 * time.Second
)

type Connector interface {
	ID() string
	Accept()
	SendMessage(data any)
	SetOnConnectHandler(fn func(c Connector))
	SetOnCloseHandler(fn func(c Connector))
	SetOnReceiveHandler(fn func(c Connector, payload []byte))
	SetOnPingHandler(fn func(c Connector))
}

func NewConnector(id string, ws *websocket.Conn) Connector {
	return &connector{
		id:       id,
		ws:       ws,
		receiver: make(chan []byte),
		sender:   make(chan []byte),
		closer:   make(chan bool),
		isClosed: false,
		ticker:   time.NewTicker(pingPeriod),
		live:     time.NewTicker(liveCheckPeriod),
	}
}

type connector struct {
	id       string
	ws       *websocket.Conn
	receiver chan []byte
	sender   chan []byte
	closer   chan bool
	isClosed bool
	ticker   *time.Ticker
	live     *time.Ticker

	onConnect func(connector Connector)
	onClose   func(connector Connector)
	onReceive func(connector Connector, message []byte)
	onPing    func(connector Connector)
}

func (c *connector) ID() string {
	return c.id
}

func (c *connector) SendMessage(data any) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	c.sender <- b
}

func (c *connector) Accept() {
	go c.readLoop()
	go c.runLoop()
	c.connect()
}

func (c *connector) SetOnConnectHandler(fn func(connector Connector)) {
	c.onConnect = fn
}

func (c *connector) SetOnCloseHandler(fn func(connector Connector)) {
	c.onClose = fn
}

func (c *connector) SetOnPingHandler(fn func(connector Connector)) {
	c.onPing = fn
}

func (c *connector) SetOnReceiveHandler(fn func(connector Connector, message []byte)) {
	c.onReceive = fn
}

func (c *connector) readLoop() {
	defer func() {
		select {
		case _, ok := <-c.closer:
			if ok == false {
				return
			}
			break
		default:
			break
		}
		c.closer <- true
	}()
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			log.Printf("Error: %v\n", err)
			break
		}
		c.receiver <- message
	}
}

func (c *connector) runLoop() {
	defer func() {
		c.close()
	}()

	c.ws.SetPingHandler(func(appData string) error {
		log.Println("ping:", appData)
		c.ticker.Reset(pingPeriod)
		c.live.Reset(liveCheckPeriod)
		go c.write(websocket.PongMessage, []byte{})
		go c.ping()
		return nil
	})

	for {
		select {
		case message := <-c.receiver:
			go c.receive(message)
		case message := <-c.sender:
			go c.write(websocket.TextMessage, message)
		case <-c.ticker.C:
			go c.write(websocket.PingMessage, []byte{})
		case <-c.live.C:
			return
		case <-c.closer:
			return
		}
	}
}

func panicRecover() {
	if r := recover(); r != nil {
		log.Println("receive panic:", r)
	}
}

func (c *connector) connect() {
	defer panicRecover()

	if c.onConnect != nil {
		c.onConnect(c)
	}
}

func (c *connector) close() {
	defer panicRecover()

	if c.isClosed == false {
		c.isClosed = true

		if c.onClose != nil {
			c.onClose(c)
		}

		close(c.receiver)
		close(c.sender)
		close(c.closer)

		c.live.Stop()
		c.ticker.Stop()
		_ = c.ws.Close()
	}
}

func (c *connector) receive(message []byte) {
	defer panicRecover()

	if c.onReceive != nil {
		c.onReceive(c, message)
	}
}

func (c *connector) write(messageType int, payload []byte) {
	defer panicRecover()

	if err := c.ws.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		panic(err)
	}
	if err := c.ws.WriteMessage(messageType, payload); err != nil {
		log.Printf("WriteMessage Error: %v\n", err)
		panic(err)
	}
}

func (c *connector) ping() {
	defer panicRecover()

	if c.onPing != nil {
		c.onPing(c)
	}
}
