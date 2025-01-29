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
		fmt.Println("Trailer already removed")
		return nil
	}

	// Confirm removal
	if !ConfirmEntry(trailer) {
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
