package dbFuncs

import (
	"database/sql"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func UpdateCombo(combo1, combo2 models.Combination) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE combinations SET trailer_fleet_num = ? WHERE truck_fleet_num = ?", combo1.TrailerFleetNum, combo1.TruckFleetNum)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE combinations SET trailer_fleet_num = ? WHERE truck_fleet_num = ?", combo2.TrailerFleetNum, combo2.TruckFleetNum)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTyreCondition(tyreID string, condition int) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a new entry with updated condition
	_, err = db.Exec("INSERT INTO tyres (tyreID, size, brand, model, supplier, price, position, location, state, condition, startingTread, archived) SELECT tyreID, size, brand, model, supplier, price, position, location, state, ?, startingTread, archived FROM tyres WHERE tyreID = ? ORDER BY created_at DESC LIMIT 1", condition, tyreID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTyreLocation(tyreID string, location string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a new entry with updated location
	_, err = db.Exec("INSERT INTO tyres (tyreID, size, brand, model, supplier, price, position, location, state, condition, startingTread, archived) SELECT tyreID, size, brand, model, supplier, price, position, ?, state, condition, startingTread, archived FROM tyres WHERE tyreID = ? ORDER BY created_at DESC LIMIT 1", location, tyreID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTyrePosition(tyreID string, position string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a new entry with updated position
	_, err = db.Exec("INSERT INTO tyres (tyreID, size, brand, model, supplier, price, position, location, state, condition, startingTread, archived) SELECT tyreID, size, brand, model, supplier, price, ?, location, state, condition, startingTread, archived FROM tyres WHERE tyreID = ? ORDER BY created_at DESC LIMIT 1", position, tyreID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTyreState(tyreID string, state string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a new entry with updated state
	_, err = db.Exec("INSERT INTO tyres (tyreID, size, brand, model, supplier, price, position, location, state, condition, startingTread, archived) SELECT tyreID, size, brand, model, supplier, price, position, location, ?, condition, startingTread, archived FROM tyres WHERE tyreID = ? ORDER BY created_at DESC LIMIT 1", state, tyreID)
	if err != nil {
		return err
	}

	return nil
}
