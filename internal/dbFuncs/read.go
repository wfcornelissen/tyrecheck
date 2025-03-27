package dbFuncs

import (
	"database/sql"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func ReadTyre(tyreID string) (models.Tyre, error) {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Tyre{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tyres WHERE id = ? ORDER BY rowid DESC LIMIT 1", tyreID)
	if err != nil {
		return models.Tyre{}, err
	}
	defer rows.Close()

	// Scan the rows into a tyre struct
	var tyre models.Tyre
	if rows.Next() {
		err = rows.Scan(&tyre.ID, &tyre.Size, &tyre.Brand, &tyre.Model, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived)
		if err != nil {
			return models.Tyre{}, err
		}
	}

	return tyre, nil
}
