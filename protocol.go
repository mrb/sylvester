package sylvester

var (
	Messages = map[string][]byte{
		"START": {0, 0},
		"EXIT":  {0, 1},
		"PING":  {0, 2},
		"PONG":  {0, 3},
	}
)

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
