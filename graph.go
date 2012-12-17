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

func (g *Graph) NewNode(nodeType string, nodeArgs ...string) *Node {
	node := NewNode()

	g.nodemap[node.Id()] = node

	return node
}

func (g *Graph) Id() *[]byte {
	return &g.id
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
	edge := NewEdge(anode, bnode)

	g.edgemap[edge.Id()] = edge

	return edge
}
