package models

import (
	"fmt"
	"time"
)

type Tyre struct {
	ID            int     `db:"id" sqlite:"INTEGER PRIMARY KEY AUTOINCREMENT"`
	TyreID        string  `db:"tyreID" sqlite:"TEXT NOT NULL"`
	Size          int     `db:"size" sqlite:"INTEGER"`
	Brand         string  `db:"brand" sqlite:"TEXT"`
	Model         string  `db:"model" sqlite:"TEXT"`
	Supplier      string  `db:"supplier" sqlite:"TEXT"`
	Price         float64 `db:"price" sqlite:"REAL"`
	Position      string  `db:"position" sqlite:"TEXT"`
	Location      string  `db:"location" sqlite:"TEXT"`
	State         string  `db:"state" sqlite:"TEXT"`
	Condition     int     `db:"condition" sqlite:"INTEGER"`
	StartingTread float64 `db:"startingTread" sqlite:"REAL"`
	Archived      bool    `db:"archived" sqlite:"BOOLEAN"`
	CreatedAt     string  `db:"created_at" sqlite:"TIMESTAMP"`
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

var RetreadState = map[int]string{
	1: "Retread Sent",
	2: "Retread Received",
	3: "Retread Scrapped",
}

var Rotate = map[int]string{
	1: "On rim",
	2: "With other tyre",
}

type TyreWork struct {
	ID       string    `db:"id" sqlite:"TEXT PRIMARY KEY UNIQUE"`
	TyreID   string    `db:"tyreID" sqlite:"TEXT"`
	WorkDate time.Time `db:"workDate" sqlite:"TEXT"`
	Position string    `db:"position" sqlite:"TEXT"`
	Odo      int       `db:"odo" sqlite:"INTEGER"`
}

func (t Tyre) String() string {
	return fmt.Sprintf("Tyre ID: %s\nSize: %d\nBrand: %s\nModel: %s\nSupplier: %s\nPrice: %2f\nPosition: %s\nLocation: %s\nState: %s\nCondition: %d\nStarting Tread: %f\nCreated At: %s", t.TyreID, t.Size, t.Brand, t.Model, t.Supplier, t.Price, t.Position, t.Location, t.State, t.Condition, t.StartingTread, t.CreatedAt)
}
func (t Tyre) SetPosition(position string) {
	t.Position = position
	fmt.Println("Tyre position successfully set to ", t.Position)
}
