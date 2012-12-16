```go
/*
A framework for creating networking applications that can be expressed
as a connected group of nodes and edges. All communication is done with
[]bytes.
*/

// A graph holds the nodes and edges.
g := syl.NewGraph()

// Create new nodes by adding them to the graph. They get
// registered with a graph hash map and behave according to
// the interface specified in the first argument. Optional
// arguments provide connection addresses, etc.
input, err := g.NewNode(UDP_NIONode, "localhost:2322")
if err != nil {
  return err
}

udpHandler = func(data []byte)(out []byte){
  copy(data,out)
  return out
}

input.AddHandler(udpHandler)

// The above node is a UDP Networked Input/Output node. It can receive input
// []byte data through a channel, and it listens for UDP protocol. A
// connection to _other nodes_ is done via a channel of a different type,
// and is specified by the creation of an Edge Object that connects the
// nodes.
// process is a regular IONode - it can read and write on []byte chans, and
// it usually does work instead of interfacing with the outside world.
// IONodes are stateless workhorse nodes, which should not mutate input data
// but may return data that is different in length or content.
process, err := g.NewNode(IONode)
if err != nil {
  return err
}

// output is a TCP IO Node. Note that simply declaring it as a TCP_NIONode does
// not mean that it is an automatic "TCP funnel" or the like. Sylvester is
// not very magical and prefers things be explicit. A handler function will be
// added to do the funnelling.
output, err := g.NewNode(TCP_NIONode "localhost:2323")
if err != nil {
  return err
}

// Declaring edges between nodes allows them to communicate over data channels.
// I want this to be separate from the beginning so that it is possible to
// play with what Edges are. Maybe they can enable interprocess communication,
// for example.
err = g.NewEdge(nodeone, nodetwo)
if err != nil {
  return err
}

// Right now we don't store the edge output or return a pointer. I don't want
// you to have to do anything with edges beyond declaring them. "Activating"
// the graph makes them work.
err = g.NewEdge(nodetwo, nodethree)
if err != nil {
  return err
}

// Activate the graph. Calls Read() on all nodes, arranges handlers, etc.
// Gets the data flowing.
g.Activate()
```
