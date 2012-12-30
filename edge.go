package sylvester

import ()

type Edge struct {
	id     []byte
	anode  *Node
	bnodes []*Node
}

func (e *Edge) Id() *[]byte {
	return &e.id
}

func NewEdge(anode, bnode *Node) *Edge {
	return &Edge{
		id:     newID(),
		anode:  anode,
		bnodes: []*Node{bnode},
	}
}

func NewEdges(anode *Node, bnodes []*Node) *Edge {
	return &Edge{
		id:     newID(),
		anode:  anode,
		bnodes: bnodes,
	}
}

func (e *Edge) Activate(c Channels) {
	go func() {
		for {
			select {
			case data := <-e.anode.Data:
				for _, bnode := range e.bnodes {
					bnode.Data <- data
				}
			case control := <-e.anode.Control:
				for _, bnode := range e.bnodes {
					bnode.Control <- control
				}
			case err := <-e.anode.Error:
				for _, bnode := range e.bnodes {
					bnode.Error <- err
				}
			}
		}
	}()
}
