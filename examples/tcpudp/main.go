/*

Sylvester models programs after Graphs to make it easy (hopefully) to design
and implement applications that route, aggregate, or otherwise deal with
streams of data over networks.

A Graph has nodes and edges. Nodes handle computation and edges handle
communication betwen nodes. Included in "computation" is communication
with outside datasources. Computation that a node is responsible for is
represented by "Events," which are functions that conform to a specific
interface and are attached to a node. All communication between nodes and
the outside world is done with bye slices - []byte rules.

To use this example:

- `nc -lk 127.0.0.1 2322` in one console
- `nc -luk 2323` in another
- `go build && ./tcpdup` Build and start this program
- type some stuff in the TCP console.
- Check out the UDP console. Chuckle. It breaks easily.

*/
package main

import (
	syl "github.com/mrb/sylvester"
	conn "github.com/mrb/sylvester/connections"
	"log"
	"os"
)

func main() {
	// Creates a new graph. Nodes and edges will be attached to 'graph'
	graph := syl.NewGraph()

	// The first Node is named input. Following the first node is the declaration
	// of the first event function we're seeing - it connects to a UDP address,
	// and then starts reading in a loop.
	input := graph.NewNode()
	UDPbyteReader := func(dc syl.DataChan, ec syl.ErrorChan) {
		tcp, err := conn.TcpConnect("localhost:2322")
		if err != nil {
			ec <- err
			return
		}

		data := make([]byte, 512)

		for {
			log.Print("Reading from TCP...")
			dlen, err := tcp.Read(data)
			if err != nil {
				ec <- err
				return
			}
			log.Printf("...read %d bytes from TCP", dlen)

			dc <- data[0:dlen]
		}
	}

	// Here, we're attaching the above function to the Event. This event will get
	// executed when the graph is activated.
	_ = input.NewEvent(UDPbyteReader)

	// This is a simple graph with two nodes and one edge. Here's the second node.
	output := graph.NewNode()

	// Note that TCPbyteWriter and UDPbyteReader have the same Function Type (they take
	// and return the same type of parameters). This is because all event functions must
	// conform to this signature, which gives event functions access to the Data and Event
	// channels internal to the nodes. This allows developers using Sylvester to easily
	// separate the work of an application from it's boilerplate. Events are the heart of
	// a graph.
	TCPbyteWriter := func(dc syl.DataChan, ec syl.ErrorChan) {
		udp, err := conn.UdpConnect("localhost:2323")
		if err != nil {
			ec <- err
		}

		for {
			select {
			case data := <-dc:
				log.Printf("Writing %d bytes to UDP", len(data))
				udp.Write(data)
			}
		}
		graph.Channels.Control <- []byte{0, 0, 0, 0}
	}
	// Attaching the TCPbytwWriter to the output node.
	_ = output.NewEvent(TCPbyteWriter)

	// One edge is necessary to connect the two nodes above. Order matters and
	// without this edge, no data would flow!
	graph.NewEdge(input, output)

	// "Activate" means to start data flow and event running - where the fun starts.
	graph.Activate()

	// The graph's error channel is connected to each node's error channel during
	// the Activate() call. Data sent directly to the graph's ExitChan or errors
	// surfaced from Nodes are handled in this select, which blocks and prevents
	// the program from terminating.

	// This error handling/retry works but it stinks.
	retries := 3
	rc := 0

	for {
		select {
		case <-graph.Channels.Control:
			log.Print("Received Exit Signal, exiting")
			os.Exit(0)
		case err := <-graph.Channels.Errors:
			if err == conn.ErrTCPConnection {
				if rc < retries {
					log.Print(err)
					rc += 1
					log.Print("Connection failed - Retry count: ", rc)
					graph.Activate()
				} else {
					log.Print("Retries exceeded, exiting")
					os.Exit(3)
				}
			} else {
				log.Print(err)
				os.Exit(3)
			}
		}
	}
}
