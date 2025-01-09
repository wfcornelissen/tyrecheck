package models

import "fmt"

type Tyre struct {
	ID        string
	Size      int
	Type      string
	Position  int
	Condition percentage
	Status    string
	Price     string
	Supplier  string
}

func (t *Tyre) IsValidSize(size int) bool {
	return size == 315 || size == 385
}

func (t *Tyre) IsValidType(Type string) bool {
	return t.Type == "Roll" || t.Type == "Steer" || t.Type == "Pull"
}

func (t *Tyre) IsValidPosition(Position int) bool {
	return Position > 0 && Position < 26
}

type percentage int

func (p *percentage) String() string {
	return fmt.Sprintf("%d%%", *p)
}

func (t *Tyre) isValidCondition(Condition percentage) bool {
	return Condition > 0 && Condition < 100
}
