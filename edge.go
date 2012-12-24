package sylvester

import ()

type Edge struct {
	id    []byte
	anode *Node
	bnodes []*Node
}

func (e *Edge) Id() *[]byte {
	return &e.id
}

func NewEdge(anode *Node, bnodes []*Node) *Edge {
	return &Edge{
		id:    newID(),
		anode: anode,
		bnodes: bnodes,
	}
}

func (e *Edge) Activate() {
	go func() {
		for {
			select {
			case data := <-e.anode.dataChan:
        for _, bnode := range e.bnodes {
				  bnode.dataChan <- data
        }
			}
		}
	}()
}
