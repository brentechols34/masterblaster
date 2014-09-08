package physics

import (
	"math"
	"time"
)

const (
	SimUpdatesPerSecond = 50.0
	SimUpdateInverse    = 1000.0 / SimUpdatesPerSecond
	FullCircle          = math.Pi * 2
)

type Manager struct {
	Bodies     map[uint32]*Base // List of all bodies in sim
	AddBody    chan Base        // Add new physics bodies to sim
	RemoveBody chan Base        // Remove physics bodies from sim
	Events     chan Event       // Outgoing updates
}

func (m *Manager) Run() {
	lastUpdate := time.Now()
	// Wait
	for {
		timeout := lastUpdate.Add(time.Millisecond * SimUpdateInverse).Sub(time.Now())
		wait_for_timeout := true
		for wait_for_timeout {
			select {
			case <-time.After(timeout):
				wait_for_timeout = false
				break
			case body := <-m.AddBody:
				m.Bodies[body.Id] = &body
			}
		}
		lastUpdate = time.Now()
		m.Tick(true)
	}
}

func (m *Manager) Tick(sendUpdate bool) {
	update := Event{}
	changed := false
	updateBodies := make(map[uint32]bool, len(m.Bodies))
	for _, body := range m.Bodies {
		update.Type = Move
		changed = false
		body.Velocity.Add(MultVect2(body.Acceleration, SimUpdateInverse))

		if body.Velocity.X != 0.0 {
			body.Position.X += body.Velocity.X / SimUpdatesPerSecond
			changed = true
		}
		if body.Velocity.Y != 0.0 {
			body.Position.Y += body.Velocity.Y / SimUpdatesPerSecond
			changed = true
		}

		// this is really really inefficient
		for _, other := range m.Bodies {
			if other.Id != body.Id {
				dist := math.Pow(float64(other.Position.X-body.Position.X), 2) + math.Pow(float64(other.Position.Y-body.Position.Y), 2)
				if dist <= math.Pow(float64(other.Size+body.Size), 2) {
					update.Type = Collision
					update.OtherId = other.Id
					break
				}
			}
		}

		updateBodies[body.Id] = true
		if changed && sendUpdate {
			update.Base = *body
			m.Events <- update
		}
	}
}
