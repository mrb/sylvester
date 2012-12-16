## Sylvester


### Sylvester is:

* An attempt to sketch a Go framework for network applications.
* An abstraction for a certain class of applications.

### Kobra as a Graph:

```
[AMQP] -->
inputNode (AMQP_Node - NRead() / Write()) -->
processNode(IONode - Read() / Write()) -->
outputNode (TCP_Node - Read() - NWrite()) -->
[TCP]

[In brackets] means an outside service
```

### Current Target API

```go
// Example app - Kobra

// Instantiate a graph
graph := NewGraph("kobra")

// Declare the first node, an AMQP connected node
input := graph.NewNode(syl.AMQP_NODE, "localhost:2322")

// Create a reusable passthrough handler function for use below
passthrough := func(input []byte) output []byte {
  copy(input, output)
  return output
}

// Create a handler function for AMQP messages that only sends the body
send_body := func(input {}interface) output []byte {
  output = input.(*AMQPMessage).body
  return output
}

// Attach the send_body handler to the input node
input.AddHandler(send_body)

// Declare the process node, a plain IONODE, and add the passthrough handler
process := graph.NewNode(syl.IONODE)
process.AddHandler(passthrough)

// Declare a TCP output node, and use the passs through handler
output := graph.NewNode(syl.TCP_NODE, "localhost:2323")
output.AddHandler(passthrough)

// Connect the nodes with edges
graph.NewEdge(input, process)
graph.NewEdge(process, output)

// Activate the graph, let dat flow!
graph.Activate()
```

### Thoughts:

* Need pre-parse and post-parse handler arrays for incoming network data (i.e.
  how to send just "Body" of AMQP message without having to parse flattend []byte)
* Still have to differentiate node types based on constants. How to handle the
  different functionality?
