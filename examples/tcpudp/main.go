package main

import (
	syl "github.com/mrb/sylvester"
	conn "github.com/mrb/sylvester/connections"
	"log"
	"net"
	"os"
)

func main() {
	graph := syl.NewGraph()

	input := graph.NewNode()
	input.NewAsyncEvent(tcpConnectAndRead)

	output := graph.NewNode()
	output.NewAsyncEvent(udpConnectAndWrite)

	supervisor := graph.NewNode()
	supervisor.NewAsyncEvent(makeWatcher(input, output))

	graph.NewEdge(input, output)

	graph.Activate()

	<-graph.Control
	log.Print("Received Exit Signal, exiting")
	os.Exit(0)
}

func tcpConnectAndRead(c syl.Channels) {
	tcp, tcpConnector := maketcpConnecter()
	go tcpConnector(c)

	tcpReader := maketcpReader(tcp)
	go tcpReader(c)

	for {
		select {
		case <-c.Control:
			_, tcpConnector := maketcpConnecter()
			go tcpConnector(c)
		}
	}
}

func maketcpConnecter() (tcp *net.TCPConn, event syl.Event) {
	event = func(c syl.Channels) {
		var err error
    log.Print("tcp? ", tcp)
		tcp, err = conn.TcpConnect("localhost:2323")

		if err != nil {
			c.Error <- err
		}
	}

	return tcp, event
}

func maketcpReader(tcp *net.TCPConn) (event syl.Event) {
	return func(c syl.Channels) {
		data := make([]byte, 512)

		for {
			log.Print("Reading from TCP...")
			dlen, err := tcp.Read(data)
			if err != nil {
				log.Print(err)
				c.Error <- err
				return
			}
			log.Printf("...read %d bytes from TCP", dlen)

			c.Data <- data[0:dlen]
		}
	}
}

func udpConnectAndWrite(c syl.Channels) {
	udp, udpConnector := makeudpConnecter()
	go udpConnector(c)

	udpWriter := makeudpWriter(udp)
	go udpWriter(c)

	for {
		select {}
	}
}

func makeudpConnecter() (udp *net.UDPConn, event syl.Event) {
	event = func(c syl.Channels) {
		var err error
		udp, err = conn.UdpConnect("localhost:2322")
		if err != nil {
			c.Error <- err
			return
		}
	}

	return udp, event
}

func makeudpWriter(udp *net.UDPConn) (event syl.Event) {
	return func(c syl.Channels) {
		for {
			select {
			case data := <-c.Data:
				log.Printf("Writing %d bytes to UDP", len(data))
				udp.Write(data)
			}
		}
	}
}

func makeWatcher(input *syl.Node, output *syl.Node) syl.Event {
	return func(c syl.Channels) {
		for {
			select {
			case <-input.Error:
				log.Print("IOE")
				input.Control.Next()
			case <-output.Error:
				output.Control.Next()
			}
		}
	}
}
