package physics

import (
	"github.com/brentechols34/masterblaster/GameObjects"
	"math"
	"time"
)

const (
	SimUpdatesPerSecond = 50.0
	SimUpdateSleep      = 1000.0 / SimUpdatesPerSecond
	FullCircle          = math.Pi * 2
)

type Manager struct {
	Bodies     map[uint32]GameObjects.Base
	AddBody    chan GameObjects.Base
	RemoveBody chan GameObjects.Base
	Events     chan Event
	lastUpdate time.Time
}

func (m *Manager) Run() {
	sendUpdate := 0
	// Wait
	for {
		timeout := m.lastUpdate.Add(time.Millisecond * SimUpdateSleep).Sub(time.Now())
		wait_for_timeout := true
		for wait_for_timeout {
			select {
			case <-time.After(timeout):
				wait_for_timeout = false
				break
			case body := <-m.AddBody:
				m.Bodies[body.Id] = body
			}
		}
		m.lastUpdate = time.Now()
		// m.Tick(sendUpdate%5 == 0)
		// Only set physics updates every X ticks.
		sendUpdate++
	}
}
