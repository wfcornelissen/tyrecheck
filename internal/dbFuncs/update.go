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

	_, err = db.Exec("UPDATE tyres SET condition = ? WHERE id = ?", condition, tyreID)
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

	_, err = db.Exec("UPDATE tyres SET location = ? WHERE id = ?", location, tyreID)
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

	_, err = db.Exec("UPDATE tyres SET position = ? WHERE id = ?", position, tyreID)
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

	_, err = db.Exec("UPDATE tyres SET state = ? WHERE id = ?", state, tyreID)
	if err != nil {
		return err
	}

	return nil
}
