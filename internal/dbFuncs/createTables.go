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

	tyresTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tyres (
		id TEXT PRIMARY KEY,
		size INTEGER,
		brand TEXT,
		model TEXT,
		supplier TEXT,
		price REAL,
		position TEXT,
		location TEXT,
		state TEXT,
		condition INTEGER,
		startingTread REAL,
		archived BOOLEAN
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer tyresTable.Close()

	_, err = tyresTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateTrucksTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	trucksTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS trucks (	
		id TEXT PRIMARY KEY UNIQUE,
		make TEXT,
		model TEXT,
		year INTEGER,
		registration TEXT,
		archived BOOLEAN
		)`)
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

func CreateTrailersTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	trailersTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS trailers (
		id TEXT PRIMARY KEY UNIQUE,
		make TEXT,
		model TEXT,
		year INTEGER,
		registration TEXT,
		archived BOOLEAN
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer trailersTable.Close()

	_, err = trailersTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateCombinationTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	combinationsTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS combinations (
		id TEXT PRIMARY KEY UNIQUE,
		truck_id TEXT,
		trailer_id TEXT,
		archived BOOLEAN
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer combinationsTable.Close()

	_, err = combinationsTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateTyreCheckTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	tyreCheckTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tyrechecks (
	id TEXT PRIMARY KEY UNIQUE,
	tyre_id TEXT,
	date TIMESTAMP,
	position TEXT,
	odo INTEGER
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer tyreCheckTable.Close()

	_, err = tyreCheckTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateTyreRepairTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	tyreRepairTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tyrerepairs (
		id TEXT PRIMARY KEY UNIQUE,
		tyre_id TEXT,
		date TIMESTAMP,
		odo INTEGER
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer tyreRepairTable.Close()

	_, err = tyreRepairTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateRetreadSentTable() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	retreadTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS retreadsSent (
		id TEXT PRIMARY KEY UNIQUE,
		tyre_id TEXT,
		date_sent TIMESTAMP,
		odo INTEGER
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer retreadTable.Close()

	_, err = retreadTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}
