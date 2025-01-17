package models

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
