package entries

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
)

// Finished - Set a trailer scrap status to true
func RemoveTrailer() error {
	fleetNum := ReadString("Enter fleet number: ")
	trailer, stmt, err := dbFuncs.DeleteTrailer(fleetNum)
	if err != nil {
		return err
	}

	if ConfirmEntry(trailer) {
		dbFuncs.Delete(stmt)
	}

	fmt.Println("Trailer removed successfully")

	return nil
}

// Finished
func RemoveTruck() error {
	fleetNum := ReadString("Enter fleet number: ")
	truck, stmt, err := dbFuncs.DeleteTruck(fleetNum)
	if err != nil {
		return err
	}

	if ConfirmEntry(truck) {
		dbFuncs.Delete(stmt)
	}

	fmt.Println("Truck removed successfully")

	return nil
}

func RemoveTyre() error {
	tyreID := ReadString("Please enter tyreID: ")
	tyre, stmt, err := dbFuncs.DeleteTyre(tyreID)
	if err != nil {
		return err
	}

	if ConfirmEntry(tyre) {
		dbFuncs.Delete(stmt)
	}

	fmt.Println("Tyre removed successfully")

	return nil
}
