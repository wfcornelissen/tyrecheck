package models

import (
	"fmt"
)

type Truck struct {
	FleetNumber string
	Make        string
	Model       string
	VIN         string
	Trailer     TrailerList
}

type Trailer struct {
	FleetNumber string
	Make        string
	Model       string
	VIN         string
	Tyres       TyreList
}

type TrailerList []Trailer

func (t TrailerList) Validate() error {
	const maxTrailers = 2
	if len(t) > maxTrailers {
		return fmt.Errorf("exceeded maximum number of trailers: %d", maxTrailers)
	}
	return nil
}

type TyreList []Tyre

func (t TyreList) Validate() error {
	const maxTyres = 26
	if len(t) > maxTyres {
		return fmt.Errorf("exceeded maximum number of tyres: %d", maxTyres)
	}
	return nil
}

func (t *Truck) AddTrailer(tr *Trailer) {
	// Accepts a trailer Fleetnumber, gets the trailer from the database and adds it to the TrailerList of t
	tr = GetTrailer(tr.FleetNumber)
	t.Trailer = append(t.Trailer, *tr)
}

func GetTrailer(fleetNumber string) *Trailer {
	// Gets the trailer from the database
	return &Trailer{FleetNumber: fleetNumber}
}
