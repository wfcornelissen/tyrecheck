package entries

import (
	"database/sql"
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

// Finished
func EditCondition() error {
	tyreID := ReadString("Enter the tyre ID: ")
	condition := ReadInt("Enter the new condition (mm tread depth): ")

	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	tyre.Condition = int((1 - (tyre.StartingTread - float64(condition))) * 100)

	err = dbFuncs.UpdateTyreCondition(tyre.Condition)
	if err != nil {
		return err
	}
	return nil
}

// Finished
func EditLocation(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	//Retrieves the tyre from the database and stores it in a struct variable
	var tyre models.Tyre
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Prints the current location
	fmt.Println("Current Location: ", tyre.Location)

	// Prompts the user to enter a new location
	tyre.Location = ReadString("Enter new location: ")

	// Updates the tyre in the database
	_, err = db.Exec("UPDATE tyres SET location = ? WHERE id = ?", tyre.Location, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	// Confirms the update
	fmt.Println("Tyre updated successfully")

	return nil
}

// Update to accommodate string position
func EditPosition(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	//Retrieves the tyre from the database and stores it in a struct variable
	var tyre models.Tyre
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Prints the current position
	fmt.Println("Current Position: ", tyre.Position)

	// Prompts the user to enter a new position
	tyre.SetPosition(ReadString("Enter new position: "))

	// Updates the tyre in the database
	_, err = db.Exec("UPDATE tyres SET position = ? WHERE id = ?", tyre.Position, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	// Confirms the update
	fmt.Println("Tyre updated successfully")

	return nil
}

// Finished
func EditState(tyreID string) error {
	//Receives a tyre and retrieves it from the database
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	defer db.Close()

	var tyre models.Tyre
	// Retrieves the tyre from the database and stores it in a struct variable
	err = db.QueryRow("SELECT id, condition, location, position, state FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Condition, &tyre.Location, &tyre.Position, &tyre.State)
	if err != nil {
		fmt.Println("Error retrieving tyre:", err)
		return err
	}

	// Print current and available states, then prompts the user to enter a new state
	fmt.Println("Current State: ", tyre.State)
	fmt.Println("Available states:")
	for _, state := range models.TyreState {
		fmt.Println(state)
	}
	tyre.State = ReadString("Enter new state: ")
	if tyre.State != models.TyreState[1] && tyre.State != models.TyreState[2] && tyre.State != models.TyreState[3] && tyre.State != models.TyreState[4] {
		fmt.Println("Invalid state")
		return nil
	}

	// Updates the tyre in the database
	_, err = db.Exec("UPDATE tyres SET state = ? WHERE id = ?", tyre.State, tyreID)
	if err != nil {
		fmt.Println("Error updating tyre:", err)
		return err
	}

	// Confirms the update
	fmt.Println("Tyre updated successfully")

	return nil
}
