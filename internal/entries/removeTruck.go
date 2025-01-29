package entries

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

// Finished
func RemoveTruck(fleetNum string) error {
	// Open SQLite db
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Extract truck into struct variable from db
	var truck models.Truck
	err = db.QueryRow("SELECT * FROM trucks WHERE fleet_num = ?", fleetNum).Scan(&truck.FleetNum, &truck.VIN, &truck.Reg, &truck.Make, &truck.Model, &truck.Year, &truck.Odo, &truck.Scrap)
	if err != nil {
		return err
	}

	// Check if truck is already removed
	if truck.Scrap {
		return fmt.Errorf("truck already removed")
	}

	// Confirm removal
	if ConfirmEntry(truck) {
		truck.Scrap = true
		_, err = db.Exec("UPDATE trucks SET scrap = true WHERE fleet_num = ?", fleetNum)
		if err != nil {
			return err
		}

		fmt.Println("Truck removed successfully")

		return nil
	}

	return nil
}
