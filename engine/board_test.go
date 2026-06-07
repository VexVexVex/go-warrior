package engine

import (
	"testing"

	"go-warrior/core"
)

// testUnit implements core.Unit for testing.
type testUnit struct {
	name      string
	x, y      int
	health    int
	maxHealth int
	enemy     bool
	captive   bool
	alive     bool
}

func (u *testUnit) Name() string              { return u.name }
func (u *testUnit) Position() (int, int)       { return u.x, u.y }
func (u *testUnit) SetPosition(x, y int)       { u.x, u.y = x, y }
func (u *testUnit) Health() int                { return u.health }
func (u *testUnit) MaxHealth() int             { return u.maxHealth }
func (u *testUnit) IsEnemy() bool              { return u.enemy }
func (u *testUnit) IsCaptive() bool            { return u.captive }
func (u *testUnit) TakeDamage(dmg int)         { u.health -= dmg; if u.health < 0 { u.health = 0 }; u.alive = u.health > 0 }
func (u *testUnit) IsAlive() bool              { return u.alive }
func (u *testUnit) PerformTurn(b core.Board)    {}

func newTestUnit(name string, x, y int) *testUnit {
	return &testUnit{name: name, x: x, y: y, health: 10, maxHealth: 10, alive: true}
}

// --- NewBoard ---

func TestNewBoard(t *testing.T) {
	b := NewBoard(8, 1)
	if b.Width != 8 {
		t.Errorf("Width = %d, want 8", b.Width)
	}
	if b.Height != 1 {
		t.Errorf("Height = %d, want 1", b.Height)
	}
}

// --- Place ---

func TestPlaceUnit(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("sludge", 2, 0)

	if err := b.Place(u, 2, 0); err != nil {
		t.Fatalf("Place: %v", err)
	}
	if _, exists := b.units[[2]int{2, 0}]; !exists {
		t.Error("unit not found in units map")
	}
	if _, exists := b.unitSet[u]; !exists {
		t.Error("unit not found in unitSet")
	}
	if x, y := u.Position(); x != 2 || y != 0 {
		t.Errorf("unit position = (%d, %d), want (2, 0)", x, y)
	}
}

func TestPlaceOnOccupiedCell(t *testing.T) {
	b := NewBoard(8, 1)
	b.Place(newTestUnit("a", 2, 0), 2, 0)

	err := b.Place(newTestUnit("b", 2, 0), 2, 0)
	if err == nil {
		t.Fatal("expected error placing on occupied cell")
	}
}

func TestPlaceDuplicateUnit(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("dup", 2, 0)
	b.Place(u, 2, 0)

	err := b.Place(u, 3, 0)
	if err == nil {
		t.Fatal("expected error placing duplicate unit")
	}
}

func TestPlaceNilUnit(t *testing.T) {
	b := NewBoard(8, 1)
	err := b.Place(nil, 0, 0)
	if err == nil {
		t.Fatal("expected error placing nil unit")
	}
}

func TestPlaceOnWall(t *testing.T) {
	b := NewBoard(8, 1)
	b.terrain[0][3] = '#'

	err := b.Place(newTestUnit("w", 3, 0), 3, 0)
	if err == nil {
		t.Fatal("expected error placing on wall")
	}
}

// --- Remove ---

func TestRemoveUnit(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("sludge", 2, 0)
	b.Place(u, 2, 0)

	if err := b.Remove(u); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if _, exists := b.units[[2]int{2, 0}]; exists {
		t.Error("unit still in units map after remove")
	}
	if _, exists := b.unitSet[u]; exists {
		t.Error("unit still in unitSet after remove")
	}
}

func TestRemoveUnitNotOnBoard(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("ghost", 0, 0)

	err := b.Remove(u)
	if err == nil {
		t.Fatal("expected error removing unit not on board")
	}
}

// --- Move ---

func TestMoveUnit(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("sludge", 2, 0)
	b.Place(u, 2, 0)

	if err := b.Move(u, core.East); err != nil {
		t.Fatalf("Move: %v", err)
	}
	x, y := u.Position()
	if x != 3 || y != 0 {
		t.Errorf("unit position = (%d, %d), want (3, 0)", x, y)
	}
	if _, exists := b.units[[2]int{2, 0}]; exists {
		t.Error("old position not cleared")
	}
	if _, exists := b.units[[2]int{3, 0}]; !exists {
		t.Error("unit missing at new position")
	}
}

func TestMoveUnitOutOfBounds(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("sludge", 7, 0)
	b.Place(u, 7, 0)

	err := b.Move(u, core.East)
	if err == nil {
		t.Fatal("expected error moving out of bounds")
	}
}

func TestMoveUnitIntoWall(t *testing.T) {
	b := NewBoard(8, 1)
	b.terrain[0][3] = '#'
	u := newTestUnit("sludge", 2, 0)
	b.Place(u, 2, 0)

	err := b.Move(u, core.East)
	if err == nil {
		t.Fatal("expected error moving into wall")
	}
}

func TestMoveUnitIntoOccupied(t *testing.T) {
	b := NewBoard(8, 1)
	a := newTestUnit("a", 2, 0)
	b2 := newTestUnit("b", 3, 0)
	b.Place(a, 2, 0)
	b.Place(b2, 3, 0)

	err := b.Move(a, core.East)
	if err == nil {
		t.Fatal("expected error moving into occupied cell")
	}
}

func TestMoveUnitNotOnBoard(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("ghost", 0, 0)

	err := b.Move(u, core.East)
	if err == nil {
		t.Fatal("expected error moving unit not on board")
	}
}

// --- Sense ---

func TestSenseNoDirection(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("warrior", 2, 0)
	b.Place(u, 2, 0)

	s := b.Sense(u)
	if s.Empty {
		t.Error("expected warrior cell to not be empty")
	}
}

func TestSenseAdjacentEnemy(t *testing.T) {
	b := NewBoard(8, 1)
	w := newTestUnit("warrior", 2, 0)
	e := newTestUnit("sludge", 3, 0)
	e.enemy = true
	b.Place(w, 2, 0)
	b.Place(e, 3, 0)

	s := b.Sense(w, core.East)
	if !s.Enemy {
		t.Error("expected enemy flag on adjacent cell")
	}
	if s.Unit == nil {
		t.Fatal("expected non-nil Unit pointer")
	}
	if *s.Unit != e {
		t.Error("expected Unit pointer to point to the enemy")
	}
}

func TestSenseWall(t *testing.T) {
	b := NewBoard(8, 1)
	b.terrain[0][3] = '#'
	u := newTestUnit("warrior", 2, 0)
	b.Place(u, 2, 0)

	s := b.Sense(u, core.East)
	if !s.Wall {
		t.Error("expected wall flag")
	}
}

func TestSenseEmptyCell(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("warrior", 2, 0)
	b.Place(u, 2, 0)

	s := b.Sense(u, core.East)
	if !s.Empty {
		t.Error("expected empty flag")
	}
}

func TestSenseOutOfBounds(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("warrior", 0, 0)
	b.Place(u, 0, 0)

	s := b.Sense(u, core.West)
	if !s.Wall {
		t.Error("expected wall for out-of-bounds")
	}
}

// --- UnitAt ---

func TestUnitAt(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("sludge", 2, 0)
	b.Place(u, 2, 0)

	got := b.UnitAt(2, 0)
	if got == nil {
		t.Fatal("UnitAt returned nil, want unit")
	}
	if got != u {
		t.Error("returned wrong unit")
	}
}

func TestUnitAtEmptyCell(t *testing.T) {
	b := NewBoard(8, 1)
	got := b.UnitAt(2, 0)
	if got != nil {
		t.Errorf("UnitAt returned non-nil for empty cell")
	}
}

// --- Units ---

func TestUnits(t *testing.T) {
	b := NewBoard(8, 1)
	a := newTestUnit("a", 1, 0)
	b2 := newTestUnit("b", 3, 0)
	b.Place(a, 1, 0)
	b.Place(b2, 3, 0)

	all := b.Units()
	if len(all) != 2 {
		t.Fatalf("Units() returned %d, want 2", len(all))
	}
}

func TestUnitsEmpty(t *testing.T) {
	b := NewBoard(8, 1)
	all := b.Units()
	if len(all) != 0 {
		t.Errorf("Units() returned %d, want 0", len(all))
	}
}

// --- InBounds ---

func TestInBounds(t *testing.T) {
	b := NewBoard(8, 1)

	cases := []struct {
		x, y int
		want bool
	}{
		{0, 0, true},
		{7, 0, true},
		{-1, 0, false},
		{8, 0, false},
		{0, -1, false},
		{0, 1, false},
	}

	for _, c := range cases {
		got := b.InBounds(c.x, c.y)
		if got != c.want {
			t.Errorf("InBounds(%d, %d) = %v, want %v", c.x, c.y, got, c.want)
		}
	}
}

// --- Distance ---

func TestDistance(t *testing.T) {
	b := NewBoard(8, 1)
	a := newTestUnit("a", 1, 0)
	b2 := newTestUnit("b", 5, 0)
	b.Place(a, 1, 0)
	b.Place(b2, 5, 0)

	d := b.Distance(a, b2)
	if d != 4 {
		t.Errorf("Distance = %d, want 4", d)
	}
}

func TestDistanceSameUnit(t *testing.T) {
	b := NewBoard(8, 1)
	u := newTestUnit("u", 3, 0)
	b.Place(u, 3, 0)

	d := b.Distance(u, u)
	if d != 0 {
		t.Errorf("Distance = %d, want 0", d)
	}
}
