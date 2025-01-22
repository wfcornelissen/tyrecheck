package models

type Tyre struct {
	ID            string
	Size          string
	Brand         string
	Supplier      string
	Price         float64
	Position      int
	Location      string
	State         string
	Condition     int
	StartingTread float64
}

var TyreState = map[int]string{
	1: "New",
	2: "Used",
	3: "Damaged",
	4: "Scrap",
}
