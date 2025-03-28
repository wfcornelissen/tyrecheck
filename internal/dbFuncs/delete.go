package dbFuncs

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func Delete(stmt *sql.Stmt) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func DeleteTrailer(fleetNum string) (models.Trailer, *sql.Stmt, error) {
	// Check if trailer exists
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Trailer{}, nil, err
	}
	defer db.Close()

	// Extract trailer into struct variable from db
	var trailer models.Trailer
	err = db.QueryRow("SELECT * FROM trailers WHERE fleet_num = ?", fleetNum).Scan(&trailer.FleetNum, &trailer.VIN, &trailer.Reg, &trailer.Make, &trailer.Model, &trailer.Year, &trailer.Scrap)
	if err != nil {
		fmt.Println("Trailer not found")
		return models.Trailer{}, nil, err
	}

	// Check if trailer is already removed
	if trailer.Scrap {
		return models.Trailer{}, nil, fmt.Errorf("trailer already removed")
	}

	stmt, err := db.Prepare("UPDATE trailers SET scrap = true WHERE fleet_num = ?")
	if err != nil {
		return models.Trailer{}, nil, err
	}

	return trailer, stmt, nil
}
