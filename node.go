package sylvester

import (
	"log"
)

type DChan chan []byte
type EChan chan []error
type EFunc func(DChan) (DChan, EChan)

type Node struct {
	id         []byte
	data       []byte
	epochFuncs []EFunc
	input      chan<- []byte
	output     <-chan []byte
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	return &Node{
		id:         newID(),
		data:       nil,
		epochFuncs: nil,
		input:      make(chan<- []byte, 10),
		output:     make(<-chan []byte, 10),
	}
}

type IONode interface {
	Read() (data chan<- []byte, err <-chan error)
	Write(instrunctions []func()) (data <-chan []byte, err <-chan error)
}

func (n *Node) AddEFunc(newEpochFunc EFunc) (err error) {
	n.epochFuncs = append(n.epochFuncs, newEpochFunc)
	return nil
}

func (n *Node) Activate() {
	for _, efunc := range n.epochFuncs {
		log.Printf("%s %s", n.Id(), efunc)
	}
}
