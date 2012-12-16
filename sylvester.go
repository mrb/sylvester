package sylvester

import (
	"fmt"
	"log"
	"time"
	"reflect"
)

func NewGraph() *Graph {
	return &Graph{
		id:      newID(),
		nodes:   make([]*Node, 1024),
		edges:   make([]*Edge, 1024),
		nodemap: make(map[*[]byte]*Node),
		edgemap: make(map[*[]byte]*Edge),
	}
}

func newID() []byte {
	return []byte(fmt.Sprintf("%s", time.Now().UnixNano()))
}
