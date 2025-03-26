package entries

import (
	"database/sql"
	"fmt"
	"time"
)

func CheckTyre(tyrePosition string) error {
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
