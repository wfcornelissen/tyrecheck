package entries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func AddTruck() (models.Truck, error) {
	truck := models.Truck{
		FleetNum: ReadString("Enter fleet number: "),
		VIN:      ReadString("Enter VIN: "),
		Reg:      ReadString("Enter registration: "),
		Make:     ReadString("Enter make: "),
		Model:    ReadString("Enter model: "),
		Year:     ReadInt("Enter year: "),
		Tyres:    []models.Tyre{},
	}
	if !ConfirmEntry(truck) {
		truck, err := AddTruck()
		if err != nil {
			return models.Truck{}, err
		}
		return truck, nil
	}
	if err := UploadTruckToDb(truck); err != nil {
		return models.Truck{}, err
	}
	return truck, nil
}

// Upload to SQLite db
func UploadTruckToDb(truck models.Truck) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS trucks (fleet_num TEXT, vin TEXT, reg TEXT, make TEXT, model TEXT, year INTEGER)")

	record, err := db.Prepare("INSERT INTO trucks (fleet_num, vin, reg, make, model, year) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	_, err = record.Exec(truck.FleetNum, truck.VIN, truck.Reg, truck.Make, truck.Model, truck.Year)
	if err != nil {
		return err
	}

	return nil
}
