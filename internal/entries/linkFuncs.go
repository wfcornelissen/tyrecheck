package entries

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/wfcornelissen/tyrecheck/internal/checks"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func SwopTruckTrailer() error {
	truckFleetNum1 := ReadString("Truck Fleet Number 1: ")
	combo1, err := dbFuncs.ReadCombo(truckFleetNum1)
	if err != nil {
		return err
	}

	truckFleetNum2 := ReadString("Truck Fleet Number 2: ")
	combo2, err := dbFuncs.ReadCombo(truckFleetNum2)
	if err != nil {
		return err
	}

	combo1.TrailerFleetNum, combo2.TrailerFleetNum = combo2.TrailerFleetNum, combo1.TrailerFleetNum

	err = dbFuncs.UpdateCombo(combo1, combo2)
	if err != nil {
		return err
	}

	return nil
}

func AssignTyre(fleetNum string, tyreID string) error {

	err := checks.CheckExist(fleetNum)
	if err != nil {
		return err
	}
	err = checks.CheckExist(tyreID)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var oldTyre models.Tyre
	err = db.QueryRow("SELECT * FROM tyres WHERE position = ?", fleetNum+tyreID).Scan(&oldTyre.ID, &oldTyre.Size, &oldTyre.Brand, &oldTyre.Supplier, &oldTyre.Price, &oldTyre.Position, &oldTyre.Location, &oldTyre.State, &oldTyre.Condition, &oldTyre.StartingTread, &oldTyre.Archived)
	if err != nil {
		if err == sql.ErrNoRows {
			// No old tyre exists at this position, which is fine
			fmt.Println("No existing tyre at this position, proceeding with new assignment")
		} else {
			fmt.Println("Error getting old tyre for updates:", err)
		}
	} else {
		// Only ask for location and update if there was an old tyre
		oldTyre.Location = ReadString("Enter the location for the old tyre: ")
		time.Sleep(100 * time.Millisecond)
		_, err = db.Exec("UPDATE tyres SET location =? WHERE id = ?", oldTyre.Location, oldTyre.ID)
		if err != nil {
			fmt.Println("Error updating old tyre location:", err)
		}
	}

	time.Sleep(100 * time.Millisecond)

	// Extract tyre into struct variable from db
	var tyre models.Tyre
	err = db.QueryRow("SELECT * FROM tyres WHERE id = ?", tyreID).Scan(&tyre.ID, &tyre.Size, &tyre.Brand, &tyre.Supplier, &tyre.Price, &tyre.Position, &tyre.Location, &tyre.State, &tyre.Condition, &tyre.StartingTread, &tyre.Archived)
	if err != nil {
		return err
	}

	tyre.Archived = false
	tyre.Location = "NULL"
	position := ReadInt("Enter the new position of the tyre: ")
	tyre.Position = fleetNum + strconv.Itoa(position)

	// Update tyre in db
	_, err = db.Exec("UPDATE tyres SET archived = ?, location = ?, position = ? WHERE id = ?", tyre.Archived, tyre.Location, tyre.Position, tyreID)
	if err != nil {
		return err
	}

	return nil
}

func CreateCombo() error {
	combo := models.Combination{
		TruckFleetNum:   ReadString("Truck Fleet Number: "),
		TrailerFleetNum: ReadString("Trailer Fleet Number: "),
	}

	// Call from dbFuncs
	err := dbFuncs.CreateCombinationEntry(&combo)
	if err != nil {
		return err
	}

	return nil
}
