package dbFuncs

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables() error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	err = CreateTyresTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Tyres table created")
	err = CreateTrucksTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Trucks table created")
	err = CreateTrailersTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Trailers table created")
	err = CreateCombinationTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Combinations table created")
	err = CreateTyreCheckTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Tyre checks table created")
	err = CreateTyreRepairTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Tyre repairs table created")
	err = CreateRetreadSentTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Retread sent table created")
	err = CreateRetreadReceivedTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Retread received table created")
	err = CreateRetreadScrapTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Retread scrap table created")
	err = CreateTyreRotateTable(db)
	if err != nil {
		return err
	}
	fmt.Println("Tyre rotate table created")
	return nil
}

func CreateTyresTable(db *sql.DB) error {
	tyresTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tyres (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tyreID TEXT NOT NULL,
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
		archived BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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

func CreateTrucksTable(db *sql.DB) error {
	trucksTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS trucks (	
		fleetnum TEXT PRIMARY KEY UNIQUE,
		vin TEXT,
		reg TEXT,
		make TEXT,
		model TEXT,
		year INTEGER,
		odo INTEGER,
		scrap BOOLEAN,
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

func CreateTrailersTable(db *sql.DB) error {
	trailersTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS trailers (
		fleetnum TEXT PRIMARY KEY UNIQUE,
		vin TEXT,
		reg TEXT,
		make TEXT,
		model TEXT,
		year INTEGER,
		scrap BOOLEAN,
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

func CreateCombinationTable(db *sql.DB) error {
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

func CreateTyreCheckTable(db *sql.DB) error {
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

func CreateTyreRepairTable(db *sql.DB) error {
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

func CreateRetreadSentTable(db *sql.DB) error {
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

func CreateRetreadReceivedTable(db *sql.DB) error {
	retreadReceivedTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS retreadsReceived (
		id TEXT PRIMARY KEY UNIQUE,
		tyre_id TEXT,
		date_received TIMESTAMP,
		odo INTEGER
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer retreadReceivedTable.Close()

	_, err = retreadReceivedTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateRetreadScrapTable(db *sql.DB) error {
	retreadScrapTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS retreadsScrap (
		id TEXT PRIMARY KEY UNIQUE,
		tyre_id TEXT,
		date_scraped TIMESTAMP,
		odo INTEGER
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer retreadScrapTable.Close()

	_, err = retreadScrapTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}
	return nil
}

func CreateTyreRotateTable(db *sql.DB) error {
	tyreRotateTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tyrerotates (
		id TEXT PRIMARY KEY UNIQUE,
		tyre_id TEXT,
		position TEXT,
		date_rotated TIMESTAMP,
		odo INTEGER
	)`)
	if err != nil {
		fmt.Println("Error preparing table creation:", err)
		return err
	}
	defer tyreRotateTable.Close()

	_, err = tyreRotateTable.Exec()
	if err != nil {
		fmt.Println("Error executing table creation:", err)
		return err
	}

	return nil
}
