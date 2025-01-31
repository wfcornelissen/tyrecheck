package checks

import "database/sql"

func CheckExist(id string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	err = truckExist(id)
	if err != nil {
		err = trailerExist(id)
		if err != nil {
			err = tyreExist(id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func truckExist(fleetNum string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM trucks WHERE fleet_num = ?", fleetNum)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	return nil
}

func trailerExist(fleetNum string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM trailers WHERE fleet_num = ?", fleetNum)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	return nil
}

func tyreExist(tyreID string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tyres WHERE id = ?", tyreID)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	return nil
}
