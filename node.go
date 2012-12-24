package sylvester

import ()

type DataChan chan []byte
type ErrorChan chan error
type Event func(DataChan, ErrorChan)

type Node struct {
	id        []byte
	data      []byte
	events    []Event
	dataChan  DataChan
	errorChan ErrorChan
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode() *Node {
	return &Node{
		id:        newID(),
		data:      nil,
		events:    nil,
		dataChan:  make(DataChan, 1),
		errorChan: make(ErrorChan, 1),
	}
}

func (n *Node) DataChan() DataChan {
	return n.dataChan
}

func (n *Node) NewEvent(newEvent Event) (err error) {
	n.events = append(n.events, newEvent)
	return nil
}

func (n *Node) Activate(errorChan ErrorChan) {
	// Yes, this is insane, it only handles the first handler. Fixing soon.
	go n.events[0](n.dataChan, n.errorChan)

	select {
	case err := <-n.errorChan:
		// Kick errors pushed onto the error channel back up to the Graph.
		errorChan <- err
	}
}
