package sylvester

import ()

type Event func(DataChan, ErrorChan)

type Node struct {
	id       []byte
	data     []byte
	events   []Event
	Channels *Channels
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
			Errors:  make(ErrorChan, 1),
		},
	}
}

func (n *Node) DataChan() DataChan {
	return n.Channels.Data
}

func (n *Node) NewEvent(newEvent Event) (err error) {
	n.events = append(n.events, newEvent)
	return nil
}

func (n *Node) Activate(errorChan ErrorChan) {
	// Currently only handles one Event.
	go n.events[0](n.Channels.Data, n.Channels.Errors)

	for {
		select {
		case err := <-n.Channels.Errors:
			errorChan <- err
		case _ = <-n.Channels.Control:
		}
	}
}
