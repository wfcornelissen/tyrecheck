package entries

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func ViewTrailer(fleetNum string) error {
	fmt.Println("ViewTrailer called")

	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var trailer models.Trailer
	err = db.QueryRow("SELECT * FROM trailers WHERE fleet_num = ?", fleetNum).Scan(&trailer.FleetNum, &trailer.VIN, &trailer.Reg, &trailer.Make, &trailer.Model, &trailer.Year, &trailer.Scrap)
	if err != nil {
		return err
	}

	fmt.Println(trailer)

	return nil
}

func ViewTruck(fleetNum string) error {
	fmt.Println("ViewTruck called")

	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var truck models.Truck
	err = db.QueryRow("SELECT * FROM trucks WHERE fleet_num = ?", fleetNum).Scan(&truck.FleetNum, &truck.VIN, &truck.Reg, &truck.Make, &truck.Model, &truck.Year, &truck.Odo, &truck.Scrap)
	if err != nil {
		return err
	}

	fmt.Println(truck)

	return nil
}

func ViewTyre(tyreID string) (models.Tyre, error) {
	fmt.Println("ViewTyre called")

	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return models.Tyre{}, err
	}
	defer db.Close()

	// Check if archived column exists, if not add it
	_, err = db.Exec("ALTER TABLE tyres ADD COLUMN archived BOOLEAN DEFAULT false")
	if err != nil && !strings.Contains(err.Error(), "duplicate column name") {
		return models.Tyre{}, err
	}

	var tyre models.Tyre
	err = db.QueryRow("SELECT id, size, brand, supplier, price, position, location, state, condition, startingTread, archived FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Size, &tyre.Brand, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived)
	if err != nil {
		return models.Tyre{}, err
	}

	fmt.Println(tyre)

	return tyre, nil
}
