package core

type Unit interface {
	Name() string
	Position() (x, y int)
	SetPosition(x, y int)
	Health() int
	MaxHealth() int
	IsEnemy() bool
	IsCaptive() bool
	TakeDamage(int)
	IsAlive() bool
	PerformTurn(b Board)
}
