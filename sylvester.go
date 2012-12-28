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
	Data    DataChan
	Error  ErrorChan
	Control ControlChan
}

func NewGraph() *Graph {
	return &Graph{
		id:      newID(),
		nodes:   nil,
		edges:   nil,
		nodemap: make(map[*[]byte]*Node),
		edgemap: make(map[*[]byte]*Edge),
		Channels: &Channels{
			Data:    make(DataChan, 1),
			Control: make(ControlChan, 1),
			Error:  make(ErrorChan, 1),
		},
	}
}

func newID() []byte {
	return []byte(fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(1000000)))
}
