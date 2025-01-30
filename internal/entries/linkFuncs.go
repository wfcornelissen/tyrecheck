package entries

import (
	"database/sql"
	"fmt"
)

func LinkTruckTrailer(truckFleetNum string, trailerFleetNum string) error {
	//Open db
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	//Create table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS combinations (truck_fleet_num TEXT, trailer_fleet_num TEXT)")
	if err != nil {
		return err
	}

	//Insert into table
	_, err = db.Exec("INSERT INTO combinations (truck_fleet_num, trailer_fleet_num) VALUES (?, ?)", truckFleetNum, trailerFleetNum)
	if err != nil {
		return err
	}

	return nil
}

func SwopTruckTrailer(truckFleetNum1, truckFleetNum2, trailerFleetNum1, trailerFleetNum2 string) error {
	fmt.Println("Swopping truck and trailer")
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Update combinations
	_, err = db.Exec("UPDATE combinations SET trailer_fleet_num = ? WHERE truck_fleet_num = ?", trailerFleetNum2, truckFleetNum1)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE combinations SET trailer_fleet_num = ? WHERE truck_fleet_num = ?", trailerFleetNum1, truckFleetNum2)
	if err != nil {
		return err
	}

	return nil
}
