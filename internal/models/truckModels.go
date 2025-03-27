package models

import "fmt"

type Truck struct {
	FleetNum string `db:"fleetNum" sqlite:"TEXT PRIMARY KEY"`
	VIN      string `db:"vin" sqlite:"TEXT UNIQUE"`
	Reg      string `db:"reg" sqlite:"TEXT UNIQUE"`
	Make     string `db:"make" sqlite:"TEXT"`
	Model    string `db:"model" sqlite:"TEXT"`
	Year     int    `db:"year" sqlite:"INTEGER"`
	Odo      int    `db:"odo" sqlite:"INTEGER"`
	Scrap    bool   `db:"scrap" sqlite:"BOOLEAN"`
	Tyres    []Tyre // Tyres are not stored in the trucks table, but are stored in the tyres table. Need functionality to link the two. Maybe foreign key?
	Archived bool   `db:"archived" sqlite:"BOOLEAN"`
}

type Trailer struct {
	FleetNum string `db:"fleetNum" sqlite:"TEXT PRIMARY KEY"`
	VIN      string `db:"vin" sqlite:"TEXT UNIQUE"`
	Reg      string `db:"reg" sqlite:"TEXT UNIQUE"`
	Make     string `db:"make" sqlite:"TEXT"`
	Model    string `db:"model" sqlite:"TEXT"`
	Year     int    `db:"year" sqlite:"INTEGER"`
	Scrap    bool   `db:"scrap" sqlite:"BOOLEAN"`
	Tyres    []Tyre // Tyres are not stored in the trucks table, but are stored in the tyres table. Need functionality to link the two. Maybe foreign key?
	Archived bool   `db:"archived" sqlite:"BOOLEAN"`
}

type Combination struct {
	TruckFleetNum   string `db:"truckFleetNum" sqlite:"TEXT"`
	TrailerFleetNum string `db:"trailerFleetNum" sqlite:"TEXT"`
}

type Fleet struct {
	Fleet []Combination
}

func (t Truck) String() string {
	return fmt.Sprintf("Fleet Number: %s\nVIN: %s\nRegistration: %s\nMake: %s\nModel: %s\nYear: %d\nOdometer: %d\nScrap: %t\n", t.FleetNum, t.VIN, t.Reg, t.Make, t.Model, t.Year, t.Odo, t.Scrap)
}

func (t Trailer) String() string {
	return fmt.Sprintf("Fleet Number: %s\nVIN: %s\nRegistration: %s\nMake: %s\nModel: %s\nYear: %d", t.FleetNum, t.VIN, t.Reg, t.Make, t.Model, t.Year)
}

func (f Fleet) String() string {
	return fmt.Sprintf("Fleet: %v", f.Fleet)
}
