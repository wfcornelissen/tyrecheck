package entries

import (
	"strconv"

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

func AssignTyre() error {
	tyreID := ReadString("Tyre ID: ")
	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	vehicleFleetNum := ReadString("Please enter vehicle fleet number: ")
	position := ReadInt("Please enter position: ")
	tyre.Location = "Fitted"
	tyre.Position = vehicleFleetNum + strconv.Itoa(position)

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
