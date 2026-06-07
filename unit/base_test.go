package unit

import (
	"testing"
	"go-warrior/core"
)

func TestNewBase(t *testing.T) {
	b := NewBase("sludge", 2, 3, 10, 10, true, false)
	if b.Name() != "sludge" {
		t.Errorf("Name = %q, want %q", b.Name(), "sludge")
	}
	if x, y := b.Position(); x != 2 || y != 3 {
		t.Errorf("Position = (%d, %d), want (2, 3)", x, y)
	}
	if b.Health() != 10 {
		t.Errorf("Health = %d, want 10", b.Health())
	}
	if b.MaxHealth() != 10 {
		t.Errorf("MaxHealth = %d, want 10", b.MaxHealth())
	}
	if !b.IsEnemy() {
		t.Error("IsEnemy = false, want true")
	}
	if b.IsCaptive() {
		t.Error("IsCaptive = true, want false")
	}
}

func TestName(t *testing.T) {
	b := NewBase("archer", 0, 0, 10, 10, true, false)
	if b.Name() != "archer" {
		t.Errorf("Name = %q, want %q", b.Name(), "archer")
	}
}

func TestPosition(t *testing.T) {
	b := NewBase("u", 5, 7, 10, 10, false, false)
	x, y := b.Position()
	if x != 5 || y != 7 {
		t.Errorf("Position = (%d, %d), want (5, 7)", x, y)
	}
}

func TestSetPosition(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.SetPosition(4, 6)
	x, y := b.Position()
	if x != 4 || y != 6 {
		t.Errorf("Position after SetPosition = (%d, %d), want (4, 6)", x, y)
	}
}

func TestHealth(t *testing.T) {
	b := NewBase("u", 0, 0, 7, 10, false, false)
	if b.Health() != 7 {
		t.Errorf("Health = %d, want 7", b.Health())
	}
}

func TestMaxHealth(t *testing.T) {
	b := NewBase("u", 0, 0, 7, 10, false, false)
	if b.MaxHealth() != 10 {
		t.Errorf("MaxHealth = %d, want 10", b.MaxHealth())
	}
}

func TestSetHealth(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.SetHealth(5)
	if b.Health() != 5 {
		t.Errorf("Health after SetHealth = %d, want 5", b.Health())
	}
}

func TestIsEnemy(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, true, false)
	if !b.IsEnemy() {
		t.Error("IsEnemy = false, want true")
	}
}

func TestIsNotEnemy(t *testing.T) {
	b := NewBase("captive", 0, 0, 10, 10, false, true)
	if b.IsEnemy() {
		t.Error("IsEnemy = true for captive, want false")
	}
}

func TestIsCaptive(t *testing.T) {
	b := NewBase("captive", 0, 0, 10, 10, false, true)
	if !b.IsCaptive() {
		t.Error("IsCaptive = false, want true")
	}
}

func TestIsNotCaptive(t *testing.T) {
	b := NewBase("sludge", 0, 0, 10, 10, true, false)
	if b.IsCaptive() {
		t.Error("IsCaptive = true for enemy, want false")
	}
}

func TestTakeDamage(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.TakeDamage(4)
	if b.Health() != 6 {
		t.Errorf("Health after TakeDamage(4) = %d, want 6", b.Health())
	}
}

func TestTakeDamageClampToZero(t *testing.T) {
	b := NewBase("u", 0, 0, 5, 10, false, false)
	b.TakeDamage(10)
	if b.Health() != 0 {
		t.Errorf("Health after TakeDamage(10) = %d, want 0", b.Health())
	}
}

func TestTakeDamageNegative(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.TakeDamage(-5)
	if b.Health() != 10 {
		t.Errorf("Health after TakeDamage(-5) = %d, want 10 (no-op)", b.Health())
	}
}

func TestTakeDamageZero(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.TakeDamage(0)
	if b.Health() != 10 {
		t.Errorf("Health after TakeDamage(0) = %d, want 10 (no-op)", b.Health())
	}
}

func TestIsAlive(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	if !b.IsAlive() {
		t.Error("IsAlive = false for full health, want true")
	}
}

func TestIsAliveAfterDamage(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.TakeDamage(7)
	if !b.IsAlive() {
		t.Error("IsAlive = false after 7 damage (3 HP left), want true")
	}
}

func TestIsDeadAtZeroHealth(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.TakeDamage(10)
	if b.IsAlive() {
		t.Error("IsAlive = true at 0 HP, want false")
	}
}

func TestIsDeadWhenOverkilled(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	b.TakeDamage(20)
	if b.IsAlive() {
		t.Error("IsAlive = true after overkill, want false")
	}
}

func TestPerformTurnNoOp(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	// Should not panic with a nil board — it's a no-op
	b.PerformTurn(nil)
}

func TestBaseSatisfiesCoreUnit(t *testing.T) {
	b := NewBase("u", 0, 0, 10, 10, false, false)
	var _ core.Unit = b
}
