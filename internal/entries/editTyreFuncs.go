package entries

import (
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
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
	return nil
}
