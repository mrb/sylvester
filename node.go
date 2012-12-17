package sylvester

import (
)

const (
	Ionode  = "ionode"
	Nionode = "nionode"
	TCP     = "tcp"
	AMQP    = "amqp"
)

type Node struct {
	id           []byte
	data         []byte
	instructions []func([]byte) []byte
	input        chan<- []byte
	output       <-chan []byte
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	return &Node{
		id:           newID(),
		data:         nil,
		instructions: nil,
		input:        make(chan<- []byte, 10),
		output:       make(<-chan []byte, 10),
	}
}

type INode interface {
	Read() (data chan<- []byte, err <-chan error)
}

type ONode interface {
	Write(instrunctions []func()) (data <-chan []byte, err <-chan error)
}

type IONode interface {
	INode
	ONode
}

type Networked interface {
	Connect() (err <-chan []error)
	NRead() (data chan<- []byte, err <-chan error)
	NWrite(instrunctions []func()) (data <-chan []byte, err <-chan error)
}

type NIONode interface {
	Networked
	INode
	ONode
}

func (n *Node) Activate() {

}
