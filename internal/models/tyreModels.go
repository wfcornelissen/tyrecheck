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
	return fmt.Sprintf("Tyre ID: %s\nSize: %s\nBrand: %s\nSupplier: %s\nPrice: %f\nPosition: %d\nLocation: %s\nState: %s\nCondition: %d\nStarting Tread: %f", t.ID, t.Size, t.Brand, t.Supplier, t.Price, t.Position, t.Location, t.State, t.Condition, t.StartingTread)
}
func (t Tyre) SetPosition(position int) {
	t.Position = position
	fmt.Println("Tyre position successfully set to ", t.Position)
}

func IsValidTruckPosition(position int) bool {
	return position >= 1 && position <= 10
}

func IsValidTrailerPosition(position int) bool {
	return position >= 11 && position <= 26
}

func GetAxleFromPosition(position int) int {
	switch position {
	case 1, 2:
		return 1
	case 3, 4, 5, 6:
		return 2
	case 7, 8, 9, 10:
		return 3
	case 11, 12, 13, 14:
		return 4
	case 15, 16, 17, 18:
		return 5
	case 19, 20, 21, 22:
		return 6
	case 23, 24, 25, 26:
		return 7
	default:
		return 0
	}
}
