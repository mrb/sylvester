package sylvester

var (
	Messages = map[string][]byte{
		"START": {0, 0},
		"EXIT":  {0, 1},
		"PING":  {0, 2},
		"PONG":  {0, 3},
		"NEXT":  {0, 4},
	}
)

// Convenient Access for Message Acess
func NodeStart() []byte {
	return Messages["START"]
}

func NodeExit() []byte {
	return Messages["EXIT"]
}

func NodePing() []byte {
	return Messages["PING"]
}

func NodePong() []byte {
	return Messages["PONG"]
}

func NodeNext() []byte {
	return Messages["NEXT"]
}

// Functions for Channels
func (p ControlChan) Start() {
	p <- Messages["START"]
}

func (p ControlChan) Exit() {
	p <- Messages["EXIT"]
}

func (p ControlChan) Ping() {
	p <- Messages["PING"]
}

func (p ControlChan) Pong() {
	p <- Messages["PONG"]
}

func (p ControlChan) Next() {
	p <- Messages["NEXT"]
}
