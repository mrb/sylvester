package sylvester

import ()

const (
	IONODE  = "ionode"
	NIONODE = "nionode"
)

type Node struct {
	id           []byte
	data         []byte
	instructions []func([]byte) []byte
	input        chan<- []byte
	output       <-chan []byte
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
