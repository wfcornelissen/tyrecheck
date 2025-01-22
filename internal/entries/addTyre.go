package entries

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

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
