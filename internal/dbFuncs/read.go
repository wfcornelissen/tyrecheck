package dbFuncs

import (
	"database/sql"
	"os"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func ReadTyreID(tyreID string) (models.Tyre, error) {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Tyre{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tyres WHERE tyreID = ? ORDER BY created_at DESC LIMIT 1", tyreID)
	if err != nil {
		return models.Tyre{}, err
	}
	defer rows.Close()

	// Scan the rows into a tyre struct
	var tyre models.Tyre
	if rows.Next() {
		err = rows.Scan(&tyre.ID, &tyre.TyreID, &tyre.Size, &tyre.Brand, &tyre.Model, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived, &tyre.CreatedAt)
		if err != nil {
			return models.Tyre{}, err
		}
	}

	return tyre, nil
}

func ReadTyrePos(tyrePos string) (models.Tyre, error) {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Tyre{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tyres WHERE position = ? ORDER BY created_at DESC LIMIT 1", tyrePos)
	if err != nil {
		return models.Tyre{}, err
	}
	defer rows.Close()

	var tyre models.Tyre
	if rows.Next() {
		err = rows.Scan(&tyre.ID, &tyre.TyreID, &tyre.Size, &tyre.Brand, &tyre.Model, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived, &tyre.CreatedAt)
		if err != nil {
			return models.Tyre{}, err
		}
	}

	return tyre, nil
}

func ReadTruckID(fleetNum string) (models.Truck, error) {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Truck{}, err
	}
	defer db.Close()

	var truck models.Truck
	err = db.QueryRow("SELECT * FROM trucks WHERE fleet_num = ?", fleetNum).Scan(&truck.FleetNum, &truck.VIN, &truck.Reg, &truck.Make, &truck.Model, &truck.Year, &truck.Odo, &truck.Scrap)
	if err != nil {
		return models.Truck{}, err
	}

	return truck, nil
}

func ReadCombo(truckFleetNum string) (models.Combination, error) {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Combination{}, err
	}
	defer db.Close()

	var combo models.Combination
	err = db.QueryRow("SELECT * FROM combinations WHERE truck_fleet_num = ?", truckFleetNum).Scan(&combo.TruckFleetNum, &combo.TrailerFleetNum)
	if err != nil {
		return models.Combination{}, err
	}

	return combo, nil
}

// ReadTrailerID reads a trailer by its fleet number
func ReadTrailerID(fleetNum string) (*models.Trailer, error) {
	db, err := sql.Open("sqlite3", os.Getenv("TYRECHECK_DB"))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var trailer models.Trailer
	err = db.QueryRow(`
		SELECT fleetNum, VIN, reg, make, model, year, scrap, archived
		FROM trailers
		WHERE fleetNum = ? AND archived = 0
		ORDER BY id DESC
		LIMIT 1
	`, fleetNum).Scan(&trailer.FleetNum, &trailer.VIN, &trailer.Reg, &trailer.Make, &trailer.Model, &trailer.Year, &trailer.Scrap, &trailer.Archived)

	if err != nil {
		return nil, err
	}

	return &trailer, nil
}
