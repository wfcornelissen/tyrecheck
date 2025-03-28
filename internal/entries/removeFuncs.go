package entries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

// Finished - Set a trailer scrap status to true
func RemoveTrailer() error {
	fleetNum := ReadString("Enter fleet number: ")
	trailer, stmt, err := dbFuncs.DeleteTrailer(fleetNum)
	if err != nil {
		return err
	}

	if ConfirmEntry(trailer) {
		dbFuncs.Delete(stmt)
	}

	fmt.Println("Trailer removed successfully")

	return nil
}

// Finished
func RemoveTruck() error {
	fleetNum := ReadString("Enter fleet number: ")
	truck, stmt, err := dbFuncs.DeleteTruck(fleetNum)
	if err != nil {
		return err
	}

	if ConfirmEntry(truck) {
		dbFuncs.Delete(stmt)
	}

	fmt.Println("Truck removed successfully")

	return nil
}

func RemoveTyre(tyreID string) error {
	fmt.Println("Removing tyre with ID:", tyreID)

	// Open SQLite db
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Extract tyre into struct variable from db
	var tyre models.Tyre
	err = db.QueryRow("SELECT * FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Size, &tyre.Brand, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived)
	if err != nil {
		return err
	}

	// Check if tyre is already removed
	if tyre.State == "Scrap" {
		return fmt.Errorf("tyre already removed")
	}

	// Confirm removal
	if ConfirmEntry(tyre) {
		tyre.State = "Scrap"
		tyre.Archived = true
		_, err = db.Exec("UPDATE tyres SET state = 'Scrap', archived = true WHERE id = ?", tyreID)
		if err != nil {
			return err
		}
	}

	return nil
}
