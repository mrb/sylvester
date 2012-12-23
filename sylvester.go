package sylvester

import (
	"fmt"
	"time"
)

func NewGraph() *Graph {
	return &Graph{
		id:       newID(),
		nodes:    nil,
		edges:    nil,
		nodemap:  make(map[*[]byte]*Node),
		edgemap:  make(map[*[]byte]*Edge),
		ExitChan: make(chan bool, 1),
	}
}

func newID() []byte {
	return []byte(fmt.Sprintf("%s", time.Now().UnixNano()))
}
