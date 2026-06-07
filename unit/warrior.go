package unit

import "go-warrior/core"

type Warrior struct{}

func (w *Warrior) Walk(dir ...core.Direction) error      { return nil }
func (w *Warrior) Attack(dir ...core.Direction) error     { return nil }
func (w *Warrior) Rest() error                             { return nil }
func (w *Warrior) Bind(dir ...core.Direction) error        { return nil }
func (w *Warrior) Rescue(dir ...core.Direction) error      { return nil }
func (w *Warrior) Shoot(dir ...core.Direction) error       { return nil }
func (w *Warrior) Explode(dir ...core.Direction) (bool, error) { return false, nil }
func (w *Warrior) Form(ai func(*Homunculus)) error         { return nil }

func (w *Warrior) Feel(dir ...core.Direction) core.Space   { return core.Space{} }
func (w *Warrior) Health() int                              { return 0 }
func (w *Warrior) MaxHealth() int                           { return 0 }
func (w *Warrior) Distance() int                            { return 0 }
func (w *Warrior) Listen() []core.Space                     { return nil }
func (w *Warrior) Look(dir ...core.Direction) []core.Space  { return nil }
