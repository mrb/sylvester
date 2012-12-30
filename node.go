package sylvester

import ()

type Event func(Channels)

type Node struct {
	id     []byte
	data   []byte
	events []Event
	*Channels
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	nodeId := newID()

	return &Node{
		id:     nodeId,
		data:   nil,
		events: nil,
		Channels: &Channels{
			Data:    make(DataChan, 1),
			Control: make(ControlChan, 1),
			Error:   make(ErrorChan, 1),
			NodeId:  nodeId,
		},
	}
}

func (n *Node) DataChan() DataChan {
	return n.Data
}

func (n *Node) NewEvent(newEvent Event) error {
	n.events = append(n.events, newEvent)
	return nil
}

func (n *Node) Activate(c Channels) {
	// Currently only handles one Event.
	if len(n.events) > 0 {
		go n.events[0](*n.Channels)
	}
}
