package core

type Space struct {
	Empty      bool
	Wall       bool
	Stairs     bool
	Enemy      bool
	Captive    bool
	Ticking    bool
	Homunculus bool
	Unit       *Unit
}
