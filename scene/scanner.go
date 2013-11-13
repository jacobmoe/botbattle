package scene

import (
	"github.com/vova616/GarageEngine/engine"
  "fmt"
)

type Scanner struct {
	engine.BaseComponent
  Player *BotController
}

func NewScanner() *Scanner {
	return &Scanner{BaseComponent: engine.NewComponent()}
}

func (ms *Scanner) OnComponentAdd() {
}

func (ms *Scanner) OnHit(enemy *engine.GameObject, damager *DamageDealer) {
  move := enemy.Transform().WorldPosition()
  fmt.Println(move)
  fmt.Println(enemy.Name())
  fmt.Println(ms.GameObject().Tag)
}

func (ms *Scanner) OnDie(byTimer bool) {
	ms.GameObject().Destroy()
}
