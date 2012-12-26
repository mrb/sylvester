package sylvester

import ()

type Event func(DataChan, ErrorChan)

type Node struct {
	id       []byte
	data     []byte
	events   []Event
	channels *Channels
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	return &Node{
		id:     newID(),
		data:   nil,
		events: nil,
		channels: &Channels{
			data:    make(DataChan, 1),
			control: make(ControlChan, 1),
			errors:  make(ErrorChan, 1),
		},
	}
}

func (n *Node) DataChan() DataChan {
	return n.channels.data
}

func (n *Node) NewEvent(newEvent Event) (err error) {
	n.events = append(n.events, newEvent)
	return nil
}

func (n *Node) Activate(errorChan ErrorChan) {
	// Currently only handles one Event.
	go n.events[0](n.channels.data, n.channels.errors)

	for {
		select {
		case err := <-n.channels.errors:
			errorChan <- err
		case _ = <-n.channels.control:
		}
	}
}
