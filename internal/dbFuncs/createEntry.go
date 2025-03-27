package dbFuncs

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

/*
Tyre
Truck
Trailer
Combination
TyreCheck
TyreRepair
*/

func CreateTyreEntry(tyre *models.Tyre) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tyres (id, size, brand, model, supplier, price, position, location, state, condition, startingTread, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		tyre.ID,
		tyre.Size,
		tyre.Brand,
		tyre.Model,
		tyre.Supplier,
		tyre.Price,
		tyre.Position,
		tyre.Location,
		tyre.State,
		tyre.Condition,
		tyre.StartingTread,
		tyre.Archived)
	if err != nil {
		fmt.Println("Error creating tyre entry")
		return err
	}

	fmt.Println("Tyre entry created")

	return nil
}

func CreateTruckEntry(truck *models.Truck) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO trucks (fleetNum, vin, reg, make, model, year, odo, scrap, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		truck.FleetNum,
		truck.VIN,
		truck.Reg,
		truck.Make,
		truck.Model,
		truck.Year,
		truck.Odo,
		truck.Scrap,
		truck.Archived)
	if err != nil {
		fmt.Println("Error creating truck entry")
		return err
	}

	fmt.Println("Truck entry created")
	return nil
}

func CreateTrailerEntry(trailer *models.Trailer) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO trailers (fleetNum, vin, reg, make, model, year, scrap, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		trailer.FleetNum,
		trailer.VIN,
		trailer.Reg,
		trailer.Make,
		trailer.Model,
		trailer.Year,
		trailer.Scrap,
		trailer.Archived)
	if err != nil {
		fmt.Println("Error creating trailer entry")
		return err
	}

	fmt.Println("Trailer entry created")

	return nil
}

func CreateCombinationEntry(combination *models.Combination) error {
	return nil
}

func CreateTyreCheckEntry(tyreCheck *models.Tyre) error {
	return nil
}

func CreateTyreRepairEntry(tyreRepair *models.Tyre) error {
	return nil
}
