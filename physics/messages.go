package physics

type EventType uint8

// Event types
const (
	Move      EventType = 0
	Collision EventType = 1
)

type Event struct {
	Base
	Type    EventType
	OtherId uint32 // If another physics object is involved with this update.
}
