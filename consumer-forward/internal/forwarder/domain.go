package forwarder

type Usecase interface {
	ReceiveMessage(connectionID, data string) error
}

type Handler interface {
	RunLoop()
}

type Gateway interface {
	PostTo(connectionID string, data string) error
}
