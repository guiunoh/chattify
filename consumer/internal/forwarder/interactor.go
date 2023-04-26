package forwarder

func NewUsecase(gateway Gateway) Usecase {
	return &interactor{gateway}
}

type interactor struct {
	gateway Gateway
}

func (i interactor) ReceiveMessage(connectionID, data string) error {
	return i.gateway.PostTo(connectionID, data)
}
