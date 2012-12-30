package sylvester

var (
	Messages = map[string][]byte{
		"START": {0, 0},
		"EXIT":  {0, 1},
		"PING":  {0, 2},
		"PONG":  {0, 3},
	}
)

// Convenient Access for Message Acess
func Start() []byte {
	return Messages["START"]
}

func Exit() []byte {
	return Messages["EXIT"]
}

func Ping() []byte {
	return Messages["PING"]
}

func Pong() []byte {
	return Messages["PONG"]
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
