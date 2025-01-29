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
