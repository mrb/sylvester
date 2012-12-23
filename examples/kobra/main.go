/*

Sylvester uses the idea of Graphs to make it easy to create complex networking applications
(hopefully). A Graph has nodes and edges. Nodes handle computation and edges handle
communication betwen nodes. Included in "computation" is communication with outside data
sources. Computation that a node is responsible for is represented by "Events," which are
functions that conform to a specific interface and are attached to a node. All communication
between nodes and the outside world is done with bye slices - []byte rules.

I'm still working out the cleanest way to handle errors (one handler func for errors?),
and this trivial example plus some netcat magic is helping me discover a lot of edge cases.
More to come.

*/
package main

import (
	syl "github.com/mrb/sylvester"
	"log"
	"net"
)

func main() {
	// Creates a new graph. Nodes and edges will be attached to 'graph'
	graph := syl.NewGraph()

	// The first Node is named input. Following the first node is the declaration of
	// the first event function we're seeing - it connects to a UDP address, and then
	// starts reading in a loop.
	input := graph.NewNode()
	UDPbyteReader := func(dc syl.DataChan, ec syl.ErrorChan) {
		conn, err := TcpConnect("localhost:2322")
		if err != nil {
			ec <- err
		}

		data := make([]byte, 512)

		for {
			log.Print("Reading from TCP...")
			dlen, err := conn.Read(data)
			if err != nil {
				ec <- err
			}
			log.Printf("...read %d bytes from TCP", dlen)

			dc <- data[0:dlen]
		}
	}

	// Here, we're attaching the above function to the Event. This event will get executed
	// when the graph is activated.
	_ = input.AddEvent(UDPbyteReader)

	// This is a simple graph with two nodes and one edge. Here's the second node.
	output := graph.NewNode()

	// Note that TCPbyteWriter and UDPbyteReader have the same Function Type (they take
	// and return the same type of parameters). This is because all event functions must
	// conform to this signature, which gives event functions access to the Data and Event
	// channels internal to the nodes. This allows developers using Sylvester to easily
	// separate the work of an application from it's boilerplate. Events are the heart of
	// a graph.
	TCPbyteWriter := func(dc syl.DataChan, ec syl.ErrorChan) {
		conn, err := UdpConnect("localhost:2323")
		if err != nil {
			ec <- err
		}

		for {
			select {
			case data := <-dc:
				log.Printf("Writing %d bytes to UDP", len(data))
				conn.Write(data)
			}
		}
		graph.ExitChan <- true
	}
	// Attaching the TCPbytwWriter to the output node.
	_ = output.AddEvent(TCPbyteWriter)

	// One edge is necessary to connect the two nodes above. Order matters and without
	// this edge, no data would flow!
	graph.NewEdge(input, output)

	// "Activate" means to start data flow and event running - where the fun starts!
	graph.Activate()

	// The program waits on the graph to send a bool value ot the ExitChan to signal that
	// the program should exit.
	<-graph.ExitChan
	log.Print("Received Exit Signal, exiting")
}

// Connection functions
func TcpConnect(address string) (c *net.TCPConn, err error) {
	log.Printf("[TCP] Dialing %s", address)
	tcpaddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	c, err = net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		return nil, err
	}

	c.SetKeepAlive(true)

	return c, nil
}

func UdpConnect(address string) (c *net.UDPConn, err error) {
	log.Printf("[UDP] Dialing %s", address)
	udpaddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	c, err = net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		return nil, err
	}

	return c, nil
}
