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

func Retread() error {
	sendOrReceive := ReadInt("Please select from the below list:")
	for i := 0; i <= len(models.RetreadState); i++ {
		fmt.Println(i, models.RetreadState[i])
	}
	switch sendOrReceive {
	case 1:
		SendRetread()
	case 2:
		ReceiveRetread()
	case 3:
		ScrapRetread()
	}
}

func SendRetread() error {
	tyreID := ReadString("Please enter tyre ID: ")
	tyre, err := dbFuncs.ReadTyre(tyreID)
	if err != nil {
		return err
	}

	tyre.Model = tyre.Model + "Retread"
	tyre.Position = "Retread"
	tyre.State = "Used"
	tyre.Location = "SentRetread"

	dbFuncs.CreateTyreEntry(&tyre)
	time.Sleep(1 * time.Second)
	dbFuncs.CreateRetreadSentEntry(&tyre)

	return nil
}

func ReceiveRetread() error {
	return nil
}

func ScrapRetread() error {
	return nil