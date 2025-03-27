package entries

import (
	"fmt"
	"time"

	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

// Defunct in this branch
/*func CheckTyre(tyrePosition string) error {
	fmt.Println("Logging tyre check")
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS TyreChecks (tyrePosition TEXT, checkDate TEXT, checkType TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("Insert into TyreChecks (tyrePosition, checkDate, checkType) values (?, ?, ?)", tyrePosition, time.Now(), "Check")
	if err != nil {
		return err
	}

	fmt.Println("Tyre check logged")

	return nil
}
*/

func CheckTyre() error {
	tyreCheck := models.TyreCheck{
		TyreID:    ReadString("Please enter tyre ID: "),
		CheckDate: time.Now(),
		Position:  ReadString("Please enter tyre position: "),
		Odo:       ReadInt("Please enter tyre odo: "),
	}
	if ConfirmEntry(tyreCheck) {
		dbFuncs.CreateTyreCheckEntry(&tyreCheck)
	} else {
		fmt.Println("Tyre check not logged")
	}

	return nil
}
