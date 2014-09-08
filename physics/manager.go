package physics

import (
	"math"
	"time"
)

const (
	SimUpdatesPerSecond = 50.0
	SimUpdateSeconds    = 1000.0 / SimUpdatesPerSecond
	SimUpdateInverse    = 1 / SimUpdatesPerSecond
	FullCircle          = math.Pi * 2
)

func NewManager() *Manager {
	return &Manager{
		Bodies:           make(map[uint32]*Base, 100),
		CollidableBodies: make(map[uint32]*Base, 100),
		Shutdown:         make(chan bool, 1),
		AddBody:          make(chan Base, 3),
		RemoveBody:       make(chan Base, 3),
		UpdateBody:       make(chan Base, 3),
		Events:           make(chan Event, 10),
	}
}

type Manager struct {
	Bodies           map[uint32]*Base // List of all bodies in sim
	CollidableBodies map[uint32]*Base // List of all collidable bodies in sim
	Shutdown         chan bool        // Channel to send kill message
	AddBody          chan Base        // Add new physics bodies to sim
	RemoveBody       chan Base        // Remove physics bodies from sim
	UpdateBody       chan Base        // Update a body already in system
	Events           chan Event       // Outgoing updates
}

func (m *Manager) Run() {
	lastUpdate := time.Now()
	timeout := lastUpdate.Add(time.Millisecond * SimUpdateSeconds).Sub(time.Now())
	doRun := true
	for doRun {
		select {
		case <-time.After(timeout):
			lastUpdate = time.Now()
			m.Tick(true)
			timeout = lastUpdate.Add(time.Millisecond * SimUpdateSeconds).Sub(time.Now())
		case body := <-m.AddBody:
			if CompareFloat(body.Velocity.X, 0) || CompareFloat(body.Velocity.Y, 0) ||
				CompareFloat(body.Acceleration.X, 0) || CompareFloat(body.Acceleration.Y, 0) {
				body.moving = true
			}
			m.Bodies[body.Id] = &body
			m.CollidableBodies[body.Id] = &body
		case body := <-m.UpdateBody:
			if CompareFloat(body.Velocity.X, 0) || CompareFloat(body.Velocity.Y, 0) ||
				CompareFloat(body.Acceleration.X, 0) || CompareFloat(body.Acceleration.Y, 0) {
				body.moving = true
			}
			m.Bodies[body.Id] = &body
			m.CollidableBodies[body.Id] = &body
		case body := <-m.RemoveBody:
			delete(m.Bodies, body.Id)
		case <-m.Shutdown:
			doRun = false
		}
	}
}

func (m *Manager) Tick(sendUpdate bool) {
	update := Event{}
	changed := false
	updateBodies := make(map[uint32]bool, len(m.Bodies))

	for _, body := range m.Bodies {
		if body.moving {
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

			if changed {
				// this is really really inefficient
				for _, other := range m.CollidableBodies {
					if other.Id != body.Id {
						dist := math.Pow(float64(other.Position.X-body.Position.X), 2) + math.Pow(float64(other.Position.Y-body.Position.Y), 2)
						if dist <= math.Pow(float64(other.Size+body.Size), 2) {
							update.Type = Collision
							update.OtherId = other.Id
							break
						}
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
}

func doesCollide(collides chan bool, b1 Base, b2 Base) {
	dist := math.Pow(float64(b1.Position.X-b2.Position.X), 2) + math.Pow(float64(b1.Position.Y-b2.Position.Y), 2)
	if dist <= math.Pow(float64(b1.Size+b2.Size), 2) {
		collides <- true
	}
	collides <- false
}
