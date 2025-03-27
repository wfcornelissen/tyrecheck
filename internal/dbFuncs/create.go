package dbFuncs

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTyresTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

}

func CreateTrucksTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	trucksTable, err := db.Prepare("CREATE TABLE IF NOT EXISTS trucks (id TEXT, make TEXT, model TEXT, year INTEGER, registration TEXT, archived BOOLEAN)")
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer trucksTable.Close()

	_, err = trucksTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}
