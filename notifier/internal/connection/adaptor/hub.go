package adaptor

import (
	"log"
)

type Hub interface {
	Register(connector Connector)
	UnRegister(connector Connector)
	GetConnector(id string) Connector
	Broadcast(message []byte)
}

func NewHub() Hub {
	h := hub{
		connectors:     make(map[string]Connector),
		broadcastChan:  make(chan []byte),
		registerChan:   make(chan Connector),
		unregisterChan: make(chan Connector),
		getOneChan:     make(chan *result),
	}
	go h.runLoop()
	return &h
}

type hub struct {
	connectors     map[string]Connector
	registerChan   chan Connector
	unregisterChan chan Connector
	getOneChan     chan *result
	broadcastChan  chan []byte
}

type result struct {
	id     string
	result chan Connector
}

func (h *hub) runLoop() {
	for {
		select {
		case connector := <-h.registerChan:
			h.connectors[connector.ID()] = connector
			log.Printf("New connection added: %s (total connectors: %d)\n", connector.ID(), len(h.connectors))
		case connector := <-h.unregisterChan:
			if _, ok := h.connectors[connector.ID()]; ok {
				delete(h.connectors, connector.ID())
				log.Printf("Connection isClosed: %s (total connectors: %d)\n", connector.ID(), len(h.connectors))
			}
		case message := <-h.broadcastChan:
			for _, connector := range h.connectors {
				go func(c Connector, message []byte) {
					c.SendMessage(message)
				}(connector, message)
			}
		case fetch := <-h.getOneChan:
			if connector, ok := h.connectors[fetch.id]; ok {
				fetch.result <- connector
			} else {
				fetch.result <- nil
			}
		}
	}
}

func (h *hub) Register(connector Connector) {
	h.registerChan <- connector
}

func (h *hub) UnRegister(connector Connector) {
	h.unregisterChan <- connector
}

func (h *hub) GetConnector(id string) Connector {
	resultChan := make(chan Connector)
	h.getOneChan <- &result{id: id, result: resultChan}
	return <-resultChan
}

func (h *hub) Broadcast(message []byte) {
	h.broadcastChan <- message
}
