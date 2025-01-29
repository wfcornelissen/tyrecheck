package entries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func RemoveTrailer(fleetNum string) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var trailer models.Trailer
	err = db.QueryRow("SELECT * FROM trailers WHERE fleet_num = ?", fleetNum).Scan(&trailer.FleetNum, &trailer.VIN, &trailer.Reg, &trailer.Make, &trailer.Model, &trailer.Year, &trailer.Scrap)
	if err != nil {
		fmt.Println("Trailer not found")
		return err
	}

	fmt.Println(trailer)

	return nil
}
