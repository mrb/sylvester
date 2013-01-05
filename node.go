package sylvester

import (
	"bytes"
)

type Event func(Channels)

type Node struct {
	id          []byte
	data        []byte
	syncEvents  []Event
	asyncEvents []Event
	*Channels
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	nodeId := newID()

	return &Node{
		id:          nodeId,
		data:        nil,
		asyncEvents: nil,
		syncEvents:  nil,
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

func (n *Node) NewAsyncEvent(newEvent Event) error {
	n.asyncEvents = append(n.asyncEvents, newEvent)
	return nil
}

func (n *Node) NewSyncEvent(newEvent Event) error {
	n.syncEvents = append(n.syncEvents, newEvent)
	return nil
}

func (n *Node) Activate() {
	if len(n.asyncEvents) > 0 {
		n.StartAsyncEvents()
	}

	if len(n.syncEvents) > 0 {
		n.StartSyncEvents()
	}
}

func (n *Node) StartAsyncEvents() {
	for _, event := range n.asyncEvents {
		go event(*n.Channels)
	}
}

func (n *Node) StartSyncEvents() {
	for _, event := range n.syncEvents {
		go event(*n.Channels)
		select {
		case control := <-n.Control:
			if bytes.Equal(control, NodeNext()) {

      } else if bytes.Equal(control, NodeSyncEventRestart()) {
        n.StartSyncEvents()
			} else {
				n.Control <- control
			}
		}
	}
}
