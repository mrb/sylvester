package sylvester

import ()

type DChan chan []byte
type EChan chan []error
type EFunc func(DChan, EChan)

type Node struct {
	id         []byte
	data       []byte
	epochFuncs []EFunc
	dataChan   DChan
	errChan    EChan
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	return &Node{
		id:         newID(),
		data:       nil,
		epochFuncs: nil,
		dataChan:   make(DChan, 1),
		errChan:    make(EChan, 1),
	}
}

func (n *Node) DataChan() DChan {
	return n.dataChan
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
	go n.epochFuncs[0](n.dataChan, n.errChan)
}
