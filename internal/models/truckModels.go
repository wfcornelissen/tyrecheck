package models

import "fmt"

type Truck struct {
	FleetNum string
	VIN      string
	Reg      string
	Make     string
	Model    string
	Year     int
	Odo      int
	Scrap    bool
	Tyres    []Tyre
}

type Trailer struct {
	FleetNum string
	VIN      string
	Reg      string
	Make     string
	Model    string
	Year     int
	Scrap    bool
	Tyres    []Tyre
}

type Combination struct {
	TruckFleetNum   string
	TrailerFleetNum string
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
