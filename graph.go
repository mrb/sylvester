package sylvester

import ()

type Graph struct {
	id      []byte
	nodes   []*Node
	edges   []*Edge
	nodemap map[*[]byte]*Node
	edgemap map[*[]byte]*Edge
	IDGen   chan []byte
}

func (g *Graph) NewNode() *Node {
	node := &Node{
		id:           newID(),
		data:         nil,
		instructions: nil,
		input:        make(chan<- []byte, 10),
		output:       make(<-chan []byte, 10),
	}

	g.nodemap[&node.id] = node

	return node
}

func (g *Graph) Activate() {
	for _, edge := range g.edgemap {
		edge.Activate()
	}

	for _, node := range g.nodemap {
		node.Activate()
	}
}

func (g *Graph) NewEdge(anode, bnode *Node) *Edge {
	edge := &Edge{
		id:    newID(),
		anode: anode,
		bnode: bnode,
	}

	g.edgemap[&edge.id] = edge

	return edge
}
