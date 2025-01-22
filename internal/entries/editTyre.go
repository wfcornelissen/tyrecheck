package entries

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func EditCondition(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	//Retrieves the tyre from the database
	var tyre models.Tyre
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Stores the tyre in a struct variable

	fmt.Println("Current Condition: ", tyre.Condition)

	tyre.Condition = ReadInt("Enter new condition: ")

	_, err = db.Exec("UPDATE tyres SET condition = ? WHERE id = ?", tyre.Condition, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	fmt.Println("Tyre updated successfully")

	return nil
}

func EditLocation(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	//Retrieves the tyre from the database
	var tyre models.Tyre
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Stores the tyre in a struct variable

	fmt.Println("Current Location: ", tyre.Location)

	tyre.Location = ReadString("Enter new location: ")

	_, err = db.Exec("UPDATE tyres SET location = ? WHERE id = ?", tyre.Location, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	fmt.Println("Tyre updated successfully")

	return nil
}

func EditPosition(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	//Retrieves the tyre from the database
	var tyre models.Tyre
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Stores the tyre in a struct variable

	fmt.Println("Current Position: ", tyre.Position)

	tyre.Position = ReadInt("Enter new position: ")

	_, err = db.Exec("UPDATE tyres SET position = ? WHERE id = ?", tyre.Position, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	fmt.Println("Tyre updated successfully")

	return nil
}

func EditState(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	//Retrieves the tyre from the database
	var tyre models.Tyre
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Stores the tyre in a struct variable

	fmt.Println("Current State: ", tyre.State)

	tyre.State = ReadString("Enter new state: ")

	_, err = db.Exec("UPDATE tyres SET state = ? WHERE id = ?", tyre.State, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	fmt.Println("Tyre updated successfully")

	return nil
}
