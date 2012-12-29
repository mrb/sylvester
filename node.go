package sylvester

import ()

type Event func(DataChan, ControlChan, ErrorChan)

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
	return &Node{
		id:     newID(),
		data:   nil,
		events: nil,
		Channels: &Channels{
			Data:    make(DataChan, 1),
			Control: make(ControlChan, 1),
			Error:   make(ErrorChan, 1),
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

func (n *Node) Activate(errorChan ErrorChan) {
	// Currently only handles one Event.
	go n.events[0](n.Data, n.Control, n.Error)

	for {
		select {
		case err := <-n.Error:
			errorChan <- err
		case _ = <-n.Control:
		}
	}
}
