package physics

import (
	"log"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	log.Printf("Testing tick.")
	m := NewManager()

	for i := uint32(0); i < 5; i++ {
		m.Bodies[i] = &Base{
			Id:       i,
			Position: Vect2{X: float32(i * 2), Y: float32(i * 2)},
			Velocity: Vect2{X: 1, Y: 1},
			moving:   true,
		}
	}

	m.Tick(true)

	//for i := uint32(0); i < 5; i++ {
	//	event := <-m.Events
	//	log.Printf("Event: %v", event)
	//}
}

func TestTickSpeed(t *testing.T) {
	m := NewManager()

	for i := uint32(0); i < 200; i++ {
		m.Bodies[i] = &Base{
			Id:       i,
			Position: Vect2{X: float32(i * 2), Y: float32(i * 2)},
			Velocity: Vect2{X: 1, Y: 1},
			moving:   false,
		}
		m.CollidableBodies[i] = m.Bodies[i]
	}

	for i := uint32(400); i < 430; i++ {
		m.Bodies[i] = &Base{
			Id:       i,
			Position: Vect2{X: float32(i * 2), Y: float32(i * 2)},
			Velocity: Vect2{X: 1, Y: 1},
			moving:   true,
		}
		m.CollidableBodies[i] = m.Bodies[i]
	}

	for i := uint32(430); i < 730; i++ {
		m.Bodies[i] = &Base{
			Id:       i,
			Position: Vect2{X: float32(i * 2), Y: float32(i * 2)},
			Velocity: Vect2{X: 1, Y: 1},
			moving:   true,
		}
	}

	totalTickTime := int64(0)
	numTicks := 10
	for i := 0; i < numTicks; i++ {
		t := time.Now().UnixNano()
		m.Tick(false)
		totalTickTime += time.Now().UnixNano() - t
	}
	log.Printf("Tick average: %d microseconds", (totalTickTime/int64(numTicks))/1000)
}

func BenchmarkTick(b *testing.B) {
	log.Printf("Benching Tick")
	m := NewManager()

	for i := uint32(0); i < 50; i++ {
		m.Bodies[i] = &Base{
			Id:       i,
			Position: Vect2{X: float32(i * 2), Y: float32(i * 2)},
			Velocity: Vect2{X: 1, Y: 1},
			moving:   true,
		}
	}

	totalTickTime := int64(0)

	for i := 0; i < b.N; i++ {
		t := time.Now().UnixNano()
		m.Tick(false)
		totalTickTime += time.Now().UnixNano() - t
	}

	log.Printf("Tick average: %d microseconds", (totalTickTime/int64(b.N))/1000)
}
