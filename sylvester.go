package sylvester

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type DataChan chan []byte
type ErrorChan chan error
type ControlChan chan []byte

type Channels struct {
	Data    DataChan
	Error   ErrorChan
	Control ControlChan
	NodeId  []byte
}

func NewGraph() *Graph {
	return &Graph{
		id:       newID(),
		nodes:    nil,
		edges:    nil,
		nodemap:  make(map[*[]byte]*Node),
		edgemap:  make(map[*[]byte]*Edge),
		Channels: NewChannels(),
	}
}

func newID() []byte {
	return []byte(fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(1000000)))
}

func NewChannels() *Channels {
	return &Channels{
		Data:    make(DataChan, 1),
		Control: make(ControlChan, 1),
		Error:   make(ErrorChan, 1),
	}
}

func (d DataChan) Word() {
	log.Print("hmm")
}
