## Sylvester


### Sylvester is:

* An attempt to sketch a Go framework for network applications.
* An abstraction for a certain class of applications.

### Kobra as a Graph:

```
[AMQP] -->
inputNode (AMQP_Node - NRead() / Write()) -->
validatorNode(IONode - Read() / Write()) -->
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

// Attach the send_body handler to the input node
input.AddHandler(send_body)

// Declare the validator node, a plain IONODE, and add the validator handler
validator := graph.NewNode(syl.IONODE)
validator.AddHandler(metric_validator)

// Declare a TCP output node, and use the passs through handler
output := graph.NewNode(syl.TCP_NODE, "localhost:2323")
output.AddHandler(passthrough)

// Connect the nodes with edges
graph.NewEdge(input, validator)
graph.NewEdge(validator, output)

// Create a passthrough handler, copies input to output
passthrough := func(input []byte) output []byte {
  copy(input, output)
  return output
}

// Create a handler function for AMQP messages that only sends the body
send_body := func(input {}interface) output []byte {
  output = input.(*AMQPMessage).body
  return output
}

// Create a handler that checks metrics format
metric_validator := func(input []byte) output []byte{
  if isValidMetric(input){
    copy(input, output)
    return output
  } else {
    return nil
  }
}

// Activate the graph, let dat flow!
graph.Activate()
```

### Thoughts:

* Need pre-parse and post-parse handler arrays for incoming network data (i.e.
  how to send just "Body" of AMQP message without having to parse flattend []byte)
* Still have to differentiate node types based on constants. How to handle the
  different functionality?
