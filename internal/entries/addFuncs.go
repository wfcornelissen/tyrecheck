package entries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

// Finished
func AddTrailer() (models.Trailer, error) {
	// Read input
	trailer := models.Trailer{
		FleetNum: ReadString("Enter fleet number: "),
		VIN:      ReadString("Enter VIN: "),
		Reg:      ReadString("Enter registration: "),
		Make:     ReadString("Enter make: "),
		Model:    ReadString("Enter model: "),
		Year:     ReadInt("Enter year: "),
		Scrap:    false,
		Tyres:    []models.Tyre{},
	}

	// Confirm entry
	if !ConfirmEntry(trailer) {
		trailer, err := AddTrailer()
		if err != nil {
			return models.Trailer{}, err
		}
		return trailer, nil
	}

	// Upload to SQLite db
	if err := UploadTrailerToDb(trailer); err != nil {
		return models.Trailer{}, err
	}
	return trailer, nil
}

// Upload to SQLite db
func UploadTrailerToDb(trailer models.Trailer) error {
	// Open SQLite db
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	// Create table if it doesn't exist
	db.Exec("CREATE TABLE IF NOT EXISTS trailers (fleet_num TEXT, vin TEXT, reg TEXT, make TEXT, model TEXT, year INTEGER, scrap BOOLEAN)")

	// Insert trailer into db
	record, err := db.Prepare("INSERT INTO trailers (fleet_num, vin, reg, make, model, year, scrap) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	// Execute insert
	_, err = record.Exec(trailer.FleetNum, trailer.VIN, trailer.Reg, trailer.Make, trailer.Model, trailer.Year, trailer.Scrap)
	if err != nil {
		return err
	}

	fmt.Println("Trailer added successfully")
	return nil
}

func AddTruck() (models.Truck, error) {
	truck := models.Truck{
		FleetNum: ReadString("Enter fleet number: "),
		VIN:      ReadString("Enter VIN: "),
		Reg:      ReadString("Enter registration: "),
		Make:     ReadString("Enter make: "),
		Model:    ReadString("Enter model: "),
		Year:     ReadInt("Enter year: "),
		Odo:      ReadInt("Enter odometer: "),
		Scrap:    false,
		Tyres:    make([]models.Tyre, 10),
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

	db.Exec("CREATE TABLE IF NOT EXISTS trucks (fleet_num TEXT, vin TEXT, reg TEXT, make TEXT, model TEXT, year INTEGER, odo INTEGER, scrap BOOLEAN)")

	record, err := db.Prepare("INSERT INTO trucks (fleet_num, vin, reg, make, model, year, odo, scrap) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	_, err = record.Exec(truck.FleetNum, truck.VIN, truck.Reg, truck.Make, truck.Model, truck.Year, truck.Odo, truck.Scrap)
	if err != nil {
		return err
	}

	return nil
}

func AddTyre() (models.Tyre, error) {
	fmt.Println("tyre called")
	tyre := models.Tyre{
		ID:            ReadString("Tyre ID: "),
		Size:          ReadString("Tyre Size: "),
		Brand:         ReadString("Tyre Brand: "),
		Supplier:      ReadString("Tyre Supplier: "),
		Price:         ReadFloat("Tyre Price: "),
		Position:      ReadInt("Tyre Position: "),
		Location:      ReadString("Tyre Location: "),
		State:         ReadString("Tyre State: "),
		Condition:     ReadInt("Tyre Condition: "),
		StartingTread: ReadFloat("Tyre Starting Tread: "),
	}
	if !ConfirmEntry(tyre) {
		tyre, err := AddTyre()
		if err != nil {
			return models.Tyre{}, err
		}
		return tyre, nil
	}

	if err := UploadTyreToDb(tyre); err != nil {
		return models.Tyre{}, err
	}
	return tyre, nil
}

func UploadTyreToDb(tyre models.Tyre) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS tyres (id TEXT, size TEXT, brand TEXT, supplier TEXT, price REAL, position INTEGER, location TEXT, state TEXT, condition INTEGER, startingTread REAL)")

	record, err := db.Prepare("INSERT INTO tyres (id, size, brand, supplier, price, position, location, state, condition, startingTread) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	_, err = record.Exec(tyre.ID, tyre.Size, tyre.Brand, tyre.Supplier, tyre.Price, tyre.Position, tyre.Location, tyre.State, tyre.Condition, tyre.StartingTread)
	if err != nil {
		return err
	}
	return nil
}
