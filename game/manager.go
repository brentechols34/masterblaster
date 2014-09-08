package game

import (
	"github.com/brentechols34/masterblaster/GameObjects"
)

type Manager struct {
	Players map[uint32]GameObjects.Player
}

func (m *Manager) Run() {

}
