package entries

import (
	"fmt"

	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
)

func ViewTrailer() error {
	fmt.Println("ViewTrailer called")
	fleetNum := ReadString("Enter fleet number: ")
	trailer, err := dbFuncs.ReadTrailerID(fleetNum)
	if err != nil {
		fmt.Println("Error viewing trailer:", err)
	}

	fmt.Println(trailer)

	return nil
}

func ViewTruck() error {
	fmt.Println("ViewTruck called")
	fleetNum := ReadString("Enter fleet number: ")
	truck, err := dbFuncs.ReadTruckID(fleetNum)
	if err != nil {
		fmt.Println("Error viewing truck:", err)
	}

	fmt.Println(truck)

	return nil
}

func ViewTyre() error {
	fmt.Println("ViewTyre called")
	tyreID := ReadString("Enter tyre ID: ")
	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		fmt.Println("Error viewing tyre:", err)
	}

	fmt.Println(tyre)

	return nil
}
