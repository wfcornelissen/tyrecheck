package entries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

// Finished - Set a trailer scrap status to true
func RemoveTrailer(fleetNum string) error {
	// Check if trailer exists
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Extract trailer into struct variable from db
	var trailer models.Trailer
	err = db.QueryRow("SELECT * FROM trailers WHERE fleet_num = ?", fleetNum).Scan(&trailer.FleetNum, &trailer.VIN, &trailer.Reg, &trailer.Make, &trailer.Model, &trailer.Year, &trailer.Scrap)
	if err != nil {
		fmt.Println("Trailer not found")
		return err
	}

	// Check if trailer is already removed
	if trailer.Scrap {
		return fmt.Errorf("trailer already removed")
	}

	// Confirm removal
	if ConfirmEntry(trailer) {
		trailer.Scrap = true

		_, err = db.Exec("UPDATE trailers SET scrap = true WHERE fleet_num = ?", fleetNum)
		if err != nil {
			return err
		}

		fmt.Println("Trailer removed successfully")

		return nil
	}

	return fmt.Errorf("trailer not removed")
}

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
