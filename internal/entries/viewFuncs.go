package entries

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func ViewTrailer() error {
	fmt.Println("ViewTrailer called")
	fleetNum := ReadString("Enter fleet number: ")
	trailer, err := dbFuncs.ReadTyreID(fleetNum)
	if err != nil {
		fmt.Println("Error viewing trailer:", err)
	}

	fmt.Println(trailer)

	return nil
}

func ViewTruck() error {
	fmt.Println("ViewTruck called")
	fleetNum := ReadString("Enter fleet number: ")
	truck, err := dbFuncs.ReadTruckID(fleetNum)
	if err != nil {
		fmt.Println("Error viewing truck:", err)
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
