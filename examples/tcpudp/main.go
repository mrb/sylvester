package main

import (
	syl "github.com/mrb/sylvester"
	conn "github.com/mrb/sylvester/connections"
	"log"
	"os"
)

func main() {
	graph := syl.NewGraph()

	input := graph.NewNode()
	input.NewEvent(UDPbyteReader)

	output := graph.NewNode()
	output.NewEvent(TCPbyteWriter)

	errorHandler := graph.NewNode()
	errorHandler.NewEvent(ErrorHandler)

	graph.NewEdge(input, output)
	graph.NewEdge(input, errorHandler)
	graph.NewEdge(output, errorHandler)
	graph.NewEdge(errorHandler, input)
	graph.NewEdge(errorHandler, output)

	graph.Activate()

	select {
	case <-graph.Control:
		log.Print("Received Exit Signal, exiting")
		os.Exit(0)
	}
}

func Restarter(c syl.Channels) {
}

func UDPbyteReader(c syl.Channels) {
	tcp, err := conn.TcpConnect("localhost:2322")
	if err != nil {
		c.Error <- syl.NewEventError(c.NodeId, err)
		return
	}

	data := make([]byte, 512)

	for {
		log.Print("Reading from TCP...")
		dlen, err := tcp.Read(data)
		if err != nil {
			c.Error <- syl.NewEventError(c.NodeId, err)
			return
		}
		log.Printf("...read %d bytes from TCP", dlen)

		c.Data <- data[0:dlen]
	}
}

func TCPbyteWriter(c syl.Channels) {
	udp, err := conn.UdpConnect("localhost:2323")
	if err != nil {
		c.Error <- syl.NewEventError(c.NodeId, err)
	}

	for {
		select {
		case data := <-c.Data:
			log.Printf("Writing %d bytes to UDP", len(data))
			udp.Write(data)
		}
	}
}

func ErrorHandler(c syl.Channels) {
	e := 0
	r := 4
	for {
		select {
		case err := <-c.Error:
			log.Print(err)
			if err == conn.ErrTCPConnection {
				e += 1
				if e > r {
					os.Exit(3)
				}
				c.Control <- []byte{0, 0}
			}
		}
	}
}
