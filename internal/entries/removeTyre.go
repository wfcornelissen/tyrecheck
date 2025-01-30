package entries

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

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
	err = db.QueryRow("SELECT * FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Size, &tyre.Brand, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread)
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
		_, err = db.Exec("UPDATE tyres SET state = 'Scrap' WHERE id = ?", tyreID)
		if err != nil {
			return err
		}
	}

	return nil
}
