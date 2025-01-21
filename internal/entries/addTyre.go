package entries

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func AddTyre() (models.Tyre, error) {
	fmt.Println("tyre called")
	tyre := models.Tyre{
		ID:       readString("Tyre ID: "),
		Size:     readString("Tyre Size: "),
		Brand:    readString("Tyre Brand: "),
		Supplier: readString("Tyre Supplier: "),
		Price:    readFloat("Tyre Price: "),
		Position: readInt("Tyre Position: "),
		Location: readString("Tyre Location: "),
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

	db.Exec("CREATE TABLE IF NOT EXISTS tyres (id TEXT, size TEXT, brand TEXT, supplier TEXT, price REAL, position INTEGER, location TEXT)")

	record, err := db.Prepare("INSERT INTO tyres (id, size, brand, supplier, price, position, location) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer record.Close()

	_, err = record.Exec(tyre.ID, tyre.Size, tyre.Brand, tyre.Supplier, tyre.Price, tyre.Position, tyre.Location)
	if err != nil {
		return err
	}
	return nil
}
