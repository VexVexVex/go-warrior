package engine

import (
	"fmt"
	"go-warrior/core"
	"math"
)

type Board struct {
	Width, Height int
	terrain       [][]rune
	units         map[[2]int]core.Unit
	unitSet       map[core.Unit]struct{}
}

func NewBoard(width, height int) *Board {
	t := make([][]rune, height)
	for y := range t {
		t[y] = make([]rune, width)
		for x := range t[y] {
			t[y][x] = '.' //default to empty
		}
	}
	return &Board{
		Width:   width,
		Height:  height,
		terrain: t,
		units:   make(map[[2]int]core.Unit),
		unitSet: make(map[core.Unit]struct{}),
	}
}

// Place a unit at (x, y), returns error if occupied
func (b *Board) Place(unit core.Unit, x, y int) error {
	if unit == nil {
		return fmt.Errorf("nil Unit passed")
	}
	if _, exists := b.unitSet[unit]; exists {
		return fmt.Errorf("unit already on the board")
	}
	if _, occupied := b.units[[2]int{x, y}]; occupied {
		return fmt.Errorf("position already occupied")
	}
	if b.terrain[y][x] == '#' {
		return fmt.Errorf("Cannot place on wall")
	}
	b.units[[2]int{x, y}] = unit
	b.unitSet[unit] = struct{}{}
	unit.SetPosition(x, y)
	return nil
}

// Remove a unit from the board
func (b *Board) Remove(unit core.Unit) error {
	if _, exists := b.unitSet[unit]; !exists {
		return fmt.Errorf("Unit is not on the board")
	}
	delete(b.unitSet, unit)
	x, y := unit.Position()
	delete(b.units, [2]int{x, y})
	return nil
}

// Move a unit from (x,y) to (x+dx, y+dy)
func (b *Board) Move(unit core.Unit, direction core.Direction) error {
	if unit == nil {
		return fmt.Errorf("Empty unit")
	}
	if _, contains := b.unitSet[unit]; !contains {
		return fmt.Errorf("Unit not on board")
	}
	x,y := unit.Position()
	delete(b.units, [2]int{x,y})
	switch direction {
	case core.North:
		y++
	case core.South:
		y--
	case core.East:
		x++
	case core.West:
		x--
	}
	if !b.InBounds(x,y) {
		return fmt.Errorf("out of bounds movement")
	}
	if b.UnitAt(x,y) != nil {
		return fmt.Errorf("unit already there")
	}
	if b.terrain[y][x] == '#' {
		return fmt.Errorf("cannot move into wall")
	}
	unit.SetPosition(x,y)
	b.units[[2]int{x,y}] = unit
	return nil
}

// Sense: return the Space at a relative offset from a unit
func (b *Board) Sense(unit core.Unit, dir ...core.Direction) core.Space {
	x, y := unit.Position()
	if len(dir) > 0 {
		switch dir[0] {
		case core.North:
			y++
		case core.South:
			y--
		case core.East:
			x++
		case core.West:
			x--
		}
	}
	if !b.InBounds(x, y) {
		return core.Space{Wall: true}
	}
	if b.terrain[y][x] == '#' {
		return core.Space{Wall: true}
	}
	s := core.Space{Empty: true}
	if u, ok := b.units[[2]int{x, y}]; ok {
		s.Empty = false
		s.Unit = &u
		if u.IsEnemy() {
			s.Enemy = true
		}
		if u.IsCaptive() {
			s.Captive = true
		}
	}
	return s
}

// Find units at a given position
func (b *Board) UnitAt(x, y int) core.Unit {
	return b.units[[2]int{x, y}]
}

// Return all alive units on the board
func (b *Board) Units() []core.Unit {
	units := make([]core.Unit, 0)
	for u := range b.unitSet {
		units = append(units, u)
	}
	return units
}

// Check if (x, y) is within bounds
func (b *Board) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < b.Width && y < b.Height
}

// Distance between two units
func (b *Board) Distance(u1, u2 core.Unit) int {
	if !b.InBounds(u1.Position()) || !b.InBounds(u2.Position()) {
		return -1
	}
	x1, y1 := u1.Position()
	x2, y2 := u2.Position()
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
