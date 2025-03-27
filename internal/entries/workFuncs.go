package entries

import (
	"fmt"
	"time"

	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func CheckTyre() error {
	tyreCheck := models.TyreWork{
		TyreID:   ReadString("Please enter tyre ID: "),
		WorkDate: time.Now(),
		Position: ReadString("Please enter tyre position: "),
		Odo:      ReadInt("Please enter tyre odo: "),
	}
	if ConfirmEntry(tyreCheck) {
		dbFuncs.CreateTyreCheckEntry(&tyreCheck)
	} else {
		fmt.Println("Tyre check not logged")
	}

	return nil
}

func RepairTyre() error {
	tyreRepair := models.TyreWork{
		TyreID:   ReadString("Please enter tyre ID: "),
		WorkDate: time.Now(),
		Position: ReadString("Please enter tyre position: "),
		Odo:      ReadInt("Please enter tyre odo: "),
	}
	if ConfirmEntry(tyreRepair) {
		dbFuncs.CreateTyreRepairEntry(&tyreRepair)
	} else {
		fmt.Println("Tyre repair not logged")
	}

	return nil
}
