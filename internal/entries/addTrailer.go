package entries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func AddTrailer() (models.Trailer, error) {
	trailer := models.Trailer{
		FleetNum: readString("Enter fleet number: "),
		VIN:      readString("Enter VIN: "),
		Reg:      readString("Enter registration: "),
		Make:     readString("Enter make: "),
		Model:    readString("Enter model: "),
		Year:     readInt("Enter year: "),
		Tyres:    []models.Tyre{},
	}
	if !ConfirmEntry(trailer) {
		trailer, err := AddTrailer()
		if err != nil {
			return models.Trailer{}, err
		}
		return trailer, nil
	}
	if err := UploadTrailerToDb(trailer); err != nil {
		return models.Trailer{}, err
	}
	return trailer, nil
}

// Upload to SQLite db
func UploadTrailerToDb(trailer models.Trailer) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS trailers (fleet_num TEXT, vin TEXT, reg TEXT, make TEXT, model TEXT, year INTEGER)")

	record, err := db.Prepare("INSERT INTO trailers (fleet_num, vin, reg, make, model, year) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	_, err = record.Exec(trailer.FleetNum, trailer.VIN, trailer.Reg, trailer.Make, trailer.Model, trailer.Year)
	if err != nil {
		return err
	}

	return nil
}
