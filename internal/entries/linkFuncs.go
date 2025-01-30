package entries

import (
	"database/sql"
	"fmt"
	"time"
)

func ComboLink(truckFleetNum string, trailerFleetNum string) error {
	//Open db
	db, err := sql.Open("sqlite3", "./tyrecheck.db?_journal=WAL&_busy_timeout=5000")
	if err != nil {
		return err
	}
	defer db.Close()

	// Enable foreign keys and WAL mode
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return err
	}

	//Create table if it doesn't exist. Both columns are foreign keys and should be unique
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS combinations (truck_fleet_num TEXT, trailer_fleet_num TEXT, PRIMARY KEY (truck_fleet_num, trailer_fleet_num))")
	if err != nil {
		return err
	}

	// Check if the truck exist
	rows, err := db.Query("SELECT * FROM trucks WHERE fleet_num = ?", truckFleetNum)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("truck not found")
	}

	// Check if the trailer exists
	rows, err = db.Query("SELECT * FROM trailers WHERE fleet_num = ?", trailerFleetNum)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("trailer not found")
	}

	// Check if the truck already has a trailer linked
	rows, err = db.Query("SELECT trailer_fleet_num FROM combinations WHERE truck_fleet_num = ?", truckFleetNum)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		return fmt.Errorf("truck already has a trailer linked")
	}

	// Add a small delay before the insert to allow other transactions to complete
	time.Sleep(100 * time.Millisecond)

	//Insert into table
	_, err = db.Exec("INSERT INTO combinations (truck_fleet_num, trailer_fleet_num) VALUES (?, ?)", truckFleetNum, trailerFleetNum)
	if err != nil {
		return err
	}

	return nil
}

func SwopTruckTrailer(truckFleetNum1, truckFleetNum2, trailerFleetNum1, trailerFleetNum2 string) error {
	fmt.Println("Swopping truck and trailer")
	db, err := sql.Open("sqlite3", "./tyrecheck.db?_journal=WAL&_busy_timeout=5000")
	if err != nil {
		return err
	}
	defer db.Close()

	// Enable foreign keys and WAL mode
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return err
	}

	// Add a small delay before updates to allow other transactions to complete
	time.Sleep(100 * time.Millisecond)

	// Check if the old combinations exist
	rows, err := db.Query("SELECT * FROM combinations WHERE truck_fleet_num = ? AND trailer_fleet_num = ?", truckFleetNum1, trailerFleetNum1)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("old combinations not found")
	}

	rows, err = db.Query("SELECT * FROM combinations WHERE truck_fleet_num = ? AND trailer_fleet_num = ?", truckFleetNum2, trailerFleetNum2)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("old combinations not found")
	}

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
