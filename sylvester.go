package sylvester

import (
	"fmt"
	"math/rand"
	"time"
)

type DataChan chan []byte
type ErrorChan chan error
type ControlChan DataChan

type Channels struct {
	data    DataChan
	errors  ErrorChan
	control ControlChan
}

func NewGraph() *Graph {
	return &Graph{
		id:       newID(),
		nodes:    nil,
		edges:    nil,
		nodemap:  make(map[*[]byte]*Node),
		edgemap:  make(map[*[]byte]*Edge),
		ExitChan: make(chan bool, 1),
		ErrChan:  make(chan error, 1),
	}
}

func newID() []byte {
	return []byte(fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(1000000)))
}
