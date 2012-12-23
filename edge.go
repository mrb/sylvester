package sylvester

import ()

type Edge struct {
	id    []byte
	anode *Node
	bnode *Node
}

func (e *Edge) Id() *[]byte {
	return &e.id
}

func NewEdge(anode, bnode *Node) *Edge {
	return &Edge{
		id:    newID(),
		anode: anode,
		bnode: bnode,
	}
}

func (e *Edge) Activate() {
	go func() {
		for {
			select {
			case data := <-e.anode.dataChan:
				e.bnode.dataChan <- data
			}
		}
	}()
}
