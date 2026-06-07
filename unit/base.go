package unit

import "go-warrior/core"

type Base struct {
	name string
	x,y int
	health int
	maxHealth int
	isEnemy bool
	isCaptive bool
}

func NewBase(name string, x, y, health, maxHealth int, isEnemy, isCaptive bool) *Base {
	return &Base{
		name: name,
		x: x,
		y: y,
		health: health,
		maxHealth: maxHealth,
		isEnemy: isEnemy,
		isCaptive: isCaptive,
	}
}
		
func (b *Base) Name() string {
	return b.name
}

func (b *Base) Position() (x,y int) {
	return b.x,b.y
}

func (b *Base) SetPosition(x,y int) {
	b.x = x
	b.y = y
}

func (b *Base) Health() int {
	return b.health
}

func (b *Base) SetHealth(amount int) {
	b.health = amount
}

func (b *Base) MaxHealth() int {
	return b.maxHealth
}

func (b *Base) IsEnemy() bool {
	return b.isEnemy
}

func (b *Base) IsCaptive() bool {
	return b.isCaptive
}

func (b *Base) TakeDamage(amount int) {
	if amount < 1 {
		return
	}
	b.health -= amount
	if b.health < 0 {
		b.health = 0
	}
}

func (b *Base) IsAlive() bool {
	return b.health > 0
}

func (b *Base) PerformTurn(board core.Board) {
}
