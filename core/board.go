package core

type Board interface {
	Place(Unit, int, int) error
	Remove(Unit) error
}
