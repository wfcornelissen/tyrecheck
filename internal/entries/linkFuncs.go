package entries

import (
	"fmt"
	"strconv"

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

	// Get the current tyre entry
	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	// If there's an existing tyre at the target position, update its location
	existingTyre, err := dbFuncs.ReadTyrePos(fleetNum + tyreID)
	if err == nil {
		// Only ask for location and update if there was an old tyre
		existingTyre.Location = ReadString("Enter the location for the old tyre: ")
		err = dbFuncs.UpdateTyreLocation(existingTyre.TyreID, existingTyre.Location)
		if err != nil {
			fmt.Println("Error updating old tyre location:", err)
		}
	}

	// Create a new entry for the tyre with the new position
	position := ReadInt("Enter the new position of the tyre: ")
	tyre.Position = fleetNum + strconv.Itoa(position)
	tyre.Location = "NULL"
	tyre.Archived = false

	err = dbFuncs.CreateTyreEntry(&tyre)
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
