package entries

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/wfcornelissen/tyrecheck/internal/checks"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

/*
func ComboLink(truckFleetNum string, trailerFleetNum string) error {
	err := checks.CheckExist(truckFleetNum)
	if err != nil {
		return err
	}
	err = checks.CheckExist(trailerFleetNum)
	if err != nil {
		return err
	}

	//Open db
	db, err := sql.Open("sqlite3", "./tyrecheck.db?_journal=WAL&_busy_timeout=5000")
	if err != nil {
		return err
	}
	defer db.Close()

	// Enable foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return err
	}

	//Create table if it doesn't exist. Both columns are foreign keys and should be unique
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS combinations (truck_fleet_num TEXT, trailer_fleet_num TEXT, PRIMARY KEY (truck_fleet_num, trailer_fleet_num))")
	if err != nil {
		return err
	}

	// Check if the truck already has a trailer linked
	rows, err := db.Query("SELECT trailer_fleet_num FROM combinations WHERE truck_fleet_num = ?", truckFleetNum)
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
*/

func SwopTruckTrailer(truckFleetNum1, truckFleetNum2, trailerFleetNum1, trailerFleetNum2 string) error {
	err := checks.CheckExist(truckFleetNum1)
	if err != nil {
		return err
	}
	err = checks.CheckExist(truckFleetNum2)
	if err != nil {
		return err
	}
	err = checks.CheckExist(trailerFleetNum1)
	if err != nil {
		return err
	}
	err = checks.CheckExist(trailerFleetNum2)
	if err != nil {
		return err
	}

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

func CheckTruckTrailerCombo(truckFleetNum string) (models.Combination, error) {
	err := checks.CheckExist(truckFleetNum)
	if err != nil {
		return models.Combination{}, err
	}

	db, err := sql.Open("sqlite3", "./tyrecheck.db?_journal=WAL&_busy_timeout=5000")
	if err != nil {
		return models.Combination{}, err
	}
	defer db.Close()

	var combo models.Combination
	err = db.QueryRow("SELECT * FROM combinations WHERE truck_fleet_num = ?", truckFleetNum).Scan(&combo.TruckFleetNum, &combo.TrailerFleetNum)
	if err != nil {
		return models.Combination{}, err
	}

	ConfirmEntry(combo)

	return combo, nil
}

func AssignTyre(fleetNum string, tyreID string) error {

	err := checks.CheckExist(fleetNum)
	if err != nil {
		return err
	}
	err = checks.CheckExist(tyreID)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var oldTyre models.Tyre
	err = db.QueryRow("SELECT * FROM tyres WHERE position = ?", fleetNum+tyreID).Scan(&oldTyre.ID, &oldTyre.Size, &oldTyre.Brand, &oldTyre.Supplier, &oldTyre.Price, &oldTyre.Position, &oldTyre.Location, &oldTyre.State, &oldTyre.Condition, &oldTyre.StartingTread, &oldTyre.Archived)
	if err != nil {
		if err == sql.ErrNoRows {
			// No old tyre exists at this position, which is fine
			fmt.Println("No existing tyre at this position, proceeding with new assignment")
		} else {
			fmt.Println("Error getting old tyre for updates:", err)
		}
	} else {
		// Only ask for location and update if there was an old tyre
		oldTyre.Location = ReadString("Enter the location for the old tyre: ")
		time.Sleep(100 * time.Millisecond)
		_, err = db.Exec("UPDATE tyres SET location =? WHERE id = ?", oldTyre.Location, oldTyre.ID)
		if err != nil {
			fmt.Println("Error updating old tyre location:", err)
		}
	}

	time.Sleep(100 * time.Millisecond)

	// Extract tyre into struct variable from db
	var tyre models.Tyre
	err = db.QueryRow("SELECT * FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Size, &tyre.Brand, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived)
	if err != nil {
		return err
	}

	tyre.Archived = false
	tyre.Location = "NULL"
	position := ReadInt("Enter the new position of the tyre: ")
	tyre.Position = fleetNum + strconv.Itoa(position)

	// Update tyre in db
	_, err = db.Exec("UPDATE tyres SET archived = ?, location = ?, position = ? WHERE id = ?", tyre.Archived, tyre.Location, tyre.Position, tyreID)
	if err != nil {
		return err
	}

	return nil
}

func CreateCombo() error {
	combo := models.Combination{
		TruckFleetNum:   entries.ReadString("Truck Fleet Number: "),
		TrailerFleetNum: entries.ReadString("Trailer Fleet Number: "),
	}

	// Call from dbFuncs
	err := dbFuncs.CreateCombinationEntry(&combo)
	if err != nil {
		return err
	}

	return nil
}
