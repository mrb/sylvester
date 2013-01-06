package sylvester

type Event func(Channels, ControlChan)

type Node struct {
	id           []byte
	data         []byte
	syncEvents   []Event
	syncPosition int
	asyncEvents  []Event
	graph        *Graph
	*Channels
}

func (n *Node) Id() *[]byte {
	return &n.id
}

func NewNode(g *Graph) *Node {
	nodeId := newID()

	return &Node{
		id:           nodeId,
		data:         nil,
		asyncEvents:  nil,
		syncEvents:   nil,
		syncPosition: 0,
		Channels:     NewChannels(),
		graph:        g,
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
		go event(*n.Channels, n.graph.Control)
	}
}

func (n *Node) StartSyncEvents() {
	go n.syncEvents[n.syncPosition](*n.Channels, n.graph.Control)
}

func (n *Node) NextSyncEvent() {
	sp := n.syncPosition
	if sp == (len(n.syncEvents) - 1) {
		n.syncPosition = 0
	} else {
		n.syncPosition++
	}
	go n.syncEvents[n.syncPosition](*n.Channels, n.graph.Control)
}
