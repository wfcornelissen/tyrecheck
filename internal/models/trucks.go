package models

import "fmt"

type Truck struct {
	FleetNum string
	VIN      string
	Reg      string
	Make     string
	Model    string
	Year     int
	Tyres    []Tyre
}

type Trailer struct {
	FleetNum string
	VIN      string
	Reg      string
	Make     string
	Model    string
	Year     int
	Tyres    []Tyre
}

type Combination struct {
	Truck   Truck
	Trailer Trailer
}

type Fleet struct {
	Fleet []Combination
}

func (t Truck) String() string {
	return fmt.Sprintf("Fleet Number: %s\nVIN: %s\nRegistration: %s\nMake: %s\nModel: %s\nYear: %d", t.FleetNum, t.VIN, t.Reg, t.Make, t.Model, t.Year)
}

func (t Trailer) String() string {
	return fmt.Sprintf("Fleet Number: %s\nVIN: %s\nRegistration: %s\nMake: %s\nModel: %s\nYear: %d", t.FleetNum, t.VIN, t.Reg, t.Make, t.Model, t.Year)
}

func (f Fleet) String() string {
	return fmt.Sprintf("Fleet: %v", f.Fleet)
}
