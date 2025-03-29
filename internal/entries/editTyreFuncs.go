package entries

import (
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

	err = dbFuncs.UpdateTyreCondition(tyreID, tyre.Condition)
	if err != nil {
		return err
	}

	return nil
}

// Finished
func EditLocation() error {
	tyreID := ReadString("Enter the tyre ID: ")
	location := ReadString("Enter the new location: ")

	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	tyre.Location = location

	err = dbFuncs.UpdateTyreLocation(tyreID, tyre.Location)
	if err != nil {
		return err
	}

	return nil
}

// Update to accommodate string position
func EditPosition() error {
	tyreID := ReadString("Enter the tyre ID: ")
	position := ReadString("Enter the new position: ")

	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	tyre.Position = position

	err = dbFuncs.UpdateTyrePosition(tyreID, tyre.Position)
	if err != nil {
		return err
	}

	return nil
}

// Finished
func EditState() error {
	tyreID := ReadString("Enter the tyre ID: ")
	state := ReadInt("Enter the new state by selecting the number: ")
	for i := 0; i < len(models.TyreState); i++ {
		fmt.Println(i, models.TyreState[i])
	}

	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	tyre.State = models.TyreState[state]

	err = dbFuncs.UpdateTyreState(tyreID, tyre.State)
	if err != nil {
		return err
	}

	return nil
}
