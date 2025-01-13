package models

import "fmt"

type Tyre struct {
	ID        string
	Size      int
	Type      string
	Position  int
	Condition percentage
	Status    string
	Price     currency
	Supplier  string
}

func (t *Tyre) String() string {
	return fmt.Sprintf("Tyre: %s\nSize: %d\nType: %s\nPosition: %d\nCondition: %s\nStatus: %s\nPrice: %s\nSupplier: %s\n", t.ID, t.Size, t.Type, t.Position, t.Condition.String(), t.Status, t.Price.String(), t.Supplier)
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

//func (t *Tyre) isValidCondition(Condition percentage) bool {
//	return Condition > 0 && Condition < 100
//}

// func (t *Tyre) isValidStatus(Status string) bool {
// 	return Status == "Virgin" || Status == "Used" || Status == "Retreaded" || Status == "Dud"
// }

type currency float64

func (c *currency) String() string {
	return fmt.Sprintf("R%.2f", *c)
}

func (t *Tyre) IsValidPrice(Price float64) string {
	// ConfirmEntry() = CLI prompt to confirm entry of price
	if t.Type == "Steer" && (t.Price < 5500 || t.Price > 10000) {
		//ConfirmEntry() // TODO: Implement ConfirmEntry()
		return "ConfirmEntry() called here" // Placeholder for ConfirmEntry()
	}
	if t.Type == "Pull" && (t.Price < 5500 || t.Price > 8000) {
		//ConfirmEntry() // TODO: Implement ConfirmEntry()
		return "ConfirmEntry() called here" // Placeholder for ConfirmEntry()
	}
	if t.Type == "Roll" && (t.Price < 4500 || t.Price > 6000) {
		//ConfirmEntry() // TODO: Implement ConfirmEntry()
		return "ConfirmEntry() called here" // Placeholder for ConfirmEntry()
	}
	return "Invalid selections"
}

// TODO: Implement
//func (t *Tyre) isValidSupplier(Supplier string) {
// Takes input, compares to Database table of suppliers
// Returns true if valid, false if not
//}
