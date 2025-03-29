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

func DeleteTruck(fleetNum string) (models.Truck, *sql.Stmt, error) {
	// Check if truck exists
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Truck{}, nil, err
	}
	defer db.Close()

	// Extract truck into struct variable from db
	var truck models.Truck
	err = db.QueryRow("SELECT * FROM trucks WHERE fleet_num = ?", fleetNum).Scan(&truck.FleetNum, &truck.VIN, &truck.Reg, &truck.Make, &truck.Model, &truck.Year, &truck.Scrap)
	if err != nil {
		fmt.Println("Truck not found")
		return models.Truck{}, nil, err
	}

	// Check if truck is already removed
	if truck.Scrap {
		return models.Truck{}, nil, fmt.Errorf("truck already removed")
	}

	stmt, err := db.Prepare("UPDATE trucks SET scrap = true WHERE fleet_num = ?")
	if err != nil {
		return models.Truck{}, nil, err
	}

	return truck, stmt, nil
}

func DeleteTyre(tyreID string) (models.Tyre, *sql.Stmt, error) {
	// Check if tyre exists
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Tyre{}, nil, err
	}
	defer db.Close()

	// Extract tyre into struct variable from db
	var tyre models.Tyre
	err = db.QueryRow("SELECT * FROM tyres WHERE tyreID = ? ORDER BY created_at DESC LIMIT 1", tyreID).Scan(&tyre.ID, &tyre.TyreID, &tyre.Size, &tyre.Brand, &tyre.Model, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived, &tyre.CreatedAt)
	if err != nil {
		fmt.Println("Tyre not found")
		return models.Tyre{}, nil, err
	}

	if tyre.Archived {
		return models.Tyre{}, nil, fmt.Errorf("tyre already removed")
	}

	stmt, err := db.Prepare("UPDATE tyres SET archived = true WHERE tyreID = ?")
	if err != nil {
		return models.Tyre{}, nil, err
	}

	return tyre, stmt, nil
}
