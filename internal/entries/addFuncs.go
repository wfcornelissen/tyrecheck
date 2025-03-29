package entries

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func AddTrailer() (models.Trailer, error) {
	// Read input
	trailer := models.Trailer{
		FleetNum: ReadString("Enter fleet number: "),
		VIN:      ReadString("Enter VIN: "),
		Reg:      ReadString("Enter registration: "),
		Make:     ReadString("Enter make: "),
		Model:    ReadString("Enter model: "),
		Year:     ReadInt("Enter year: "),
		Scrap:    false,
	}

	// Confirm entry
	if !ConfirmEntry(trailer) {
		trailer, err := AddTrailer()
		if err != nil {
			return models.Trailer{}, err
		}
		return trailer, nil
	}

	err := dbFuncs.CreateTrailerEntry(&trailer)
	if err != nil {
		return models.Trailer{}, err
	}

	return trailer, nil
}

func AddTruck() (models.Truck, error) {
	// Create a default tyre with 0 as all values
	truck := models.Truck{
		FleetNum: ReadString("Enter fleet number: "),
		VIN:      ReadString("Enter VIN: "),
		Reg:      ReadString("Enter registration: "),
		Make:     ReadString("Enter make: "),
		Model:    ReadString("Enter model: "),
		Year:     ReadInt("Enter year: "),
		Odo:      ReadInt("Enter odometer: "),
		Scrap:    false,
	}

	if !ConfirmEntry(truck) {
		truck, err := AddTruck()
		if err != nil {
			return models.Truck{}, err
		}
		return truck, nil
	}

	err := dbFuncs.CreateTruckEntry(&truck)
	if err != nil {
		return models.Truck{}, err
	}

	return truck, nil
}

func AddTyre() (models.Tyre, error) {
	fmt.Println("tyre called")
	tyre := models.Tyre{
		ID:            ReadString("Tyre ID: "),
		Size:          ReadInt("Tyre Size: "),
		Brand:         ReadString("Tyre Brand: "),
		Model:         ReadString("Tyre Model: "),
		Supplier:      ReadString("Tyre Supplier: "),
		Price:         ReadFloat("Tyre Price: "),
		Position:      ReadString("Tyre Position: "),
		Location:      ReadString("Tyre Location: "),
		State:         ReadString("Tyre State: "),
		Condition:     ReadInt("Tyre Condition: "),
		StartingTread: ReadFloat("Tyre Starting Tread: "),
		Archived:      false,
	}
	if !ConfirmEntry(tyre) {
		tyre, err := AddTyre()
		if err != nil {
			return models.Tyre{}, err
		}
		return tyre, nil
	}

	err := dbFuncs.CreateTyreEntry(&tyre)
	if err != nil {
		if err.Error() == "tyreid already exists" {
			tyre, err := AddTyre()
			if err != nil {
				return models.Tyre{}, err
			}
			return tyre, nil
		}
		return models.Tyre{}, err
	}
	return tyre, nil
}
