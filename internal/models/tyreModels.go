package models

import (
	"fmt"
)

type Tyre struct {
	ID            string
	Size          string
	Brand         string
	Supplier      string
	Price         float64
	Position      string
	Location      string
	State         string
	Condition     int
	StartingTread float64
	Archived      bool
}

var TyreState = map[int]string{
	1: "New",
	2: "Used",
	3: "Damaged",
	4: "Scrap",
}

var TyreSize = map[int]string{
	315: "315",
	385: "385",
}

func (t Tyre) String() string {
	return fmt.Sprintf("Tyre ID: %s\nSize: %s\nBrand: %s\nSupplier: %s\nPrice: %s\nPosition: %s\nLocation: %s\nState: %s\nCondition: %d\nStarting Tread: %f", t.ID, t.Size, t.Brand, t.Supplier, t.Price, t.Position, t.Location, t.State, t.Condition, t.StartingTread)
}
func (t Tyre) SetPosition(position string) {
	t.Position = position
	fmt.Println("Tyre position successfully set to ", t.Position)
}
