package main

import (
	"fmt"
	"log"
	"time"
)

const ()

type Graph struct {
	id      []byte
	nodes   []*Node
	edges   []*Edge
	nodemap map[*[]byte]*Node
	edgemap map[*[]byte]*Edge
}

type Node struct {
	id           []byte
	data         []byte
	instructions []func([]byte) []byte
	input        chan<- []byte
	output       <-chan []byte
}

type Edge struct {
	id    []byte
	anode *Node
	bnode *Node
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

func NewGraph() *Graph {
	return &Graph{
		id:      newID(),
		nodes:   make([]*Node, 1024),
		edges:   make([]*Edge, 1024),
		nodemap: make(map[*[]byte]*Node),
		edgemap: make(map[*[]byte]*Edge),
	}
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

func (e *Edge) Activate() {
	go func() {
		for {
			select {
			case data := <-e.anode.output:
				log.Print(data)
				e.bnode.input <- data
			}
		}
	}()
}

/*
// Called by the Edge.
func (n *Node) Work() (input chan<- []byte, output <-chan []byte, err <-chan []error) {
	return input, output, err
}
*/

func (n *Node) Activate() {

}

func (n *Node) Main() {
	select {}
}

func main() {
	log.Print("starting")

	graph := NewGraph()

	input := graph.NewNode()
	process := graph.NewNode()
	output := graph.NewNode()

	graph.NewEdge(input, process)
	graph.NewEdge(process, output)

	graph.Activate()

	log.Print("finished")
}

func newID() []byte {
	return []byte(fmt.Sprintf("%s", time.Now().UnixNano()))
}
