package entries

import (
	"database/sql"
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

	return truck, nil
}

func AddTyre() (models.Tyre, error) {
	fmt.Println("tyre called")
	tyre := models.Tyre{
		ID:            ReadString("Tyre ID: "),
		Size:          ReadInt("Tyre Size: "),
		Brand:         ReadString("Tyre Brand: "),
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
		return models.Tyre{}, err
	}
	return tyre, nil
}

func UploadTyreToDb(tyre models.Tyre) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS tyres (id TEXT, size TEXT, brand TEXT, supplier TEXT, price REAL, position TEXT, location TEXT, state TEXT, condition INTEGER, startingTread REAL, archived BOOLEAN)")

	record, err := db.Prepare("INSERT INTO tyres (id, size, brand, supplier, price, position, location, state, condition, startingTread, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	_, err = record.Exec(tyre.ID, tyre.Size, tyre.Brand, tyre.Supplier, tyre.Price, tyre.Position, tyre.Location, tyre.State, tyre.Condition, tyre.StartingTread, tyre.Archived)
	if err != nil {
		return err
	}
	return nil
}
