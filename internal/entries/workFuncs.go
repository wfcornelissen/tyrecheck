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
		WorkDate: ReadDate("Please enter date work was done: "),
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
		WorkDate: ReadDate("Please enter date work was done: "),
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
		err := SendRetread()
		if err != nil {
			return err
		}
		return nil
	case 2:
		err := ReceiveRetread()
		if err != nil {
			return err
		}
		return nil
	case 3:
		err := ScrapRetread()
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("invalid option")
}

func SendRetread() error {
	tyreID := ReadString("Please enter tyre ID: ")
	tyre, err := dbFuncs.ReadTyre(tyreID)
	if err != nil {
		return err
	}

	tyre.Model = tyre.Model + "Retread"
	tyre.Position = "NULL"
	tyre.State = "Used"
	tyre.Location = "SentRetread"

	dbFuncs.CreateTyreEntry(&tyre)

	time.Sleep(1 * time.Second)

	workDate := ReadDate("Please enter work date: (dd/mm/yyyy)")
	odo := ReadInt("Please enter odo: ")
	dbFuncs.CreateRetreadSentEntry(&tyre, workDate, odo)

	return nil
}

func ReceiveRetread() error {
	tyreID := ReadString("Please enter tyre ID: ")
	tyre, err := dbFuncs.ReadTyre(tyreID)
	if err != nil {
		return err
	}

	tyre.Model = tyre.Model
	tyre.Position = "NULL"
	tyre.State = "Used"
	tyre.Location = "ReceivedRetread"

	dbFuncs.CreateTyreEntry(&tyre)

	time.Sleep(1 * time.Second)

	workDate := ReadDate("Please enter work date: (dd/mm/yyyy)")
	odo := ReadInt("Please enter odo: ")
	dbFuncs.CreateRetreadReceivedEntry(&tyre, workDate, odo)

	return nil
}

func ScrapRetread() error {
	tyreID := ReadString("Please enter tyre ID: ")
	tyre, err := dbFuncs.ReadTyre(tyreID)
	if err != nil {
		return err
	}

	tyre.Position = "Scrapped"
	tyre.State = "Scrap"
	tyre.Location = "Scrap"

	dbFuncs.CreateTyreEntry(&tyre)

	time.Sleep(1 * time.Second)

	workDate := ReadDate("Please enter work date: (dd/mm/yyyy)")
	odo := ReadInt("Please enter odo: ")
	dbFuncs.CreateRetreadScrapEntry(&tyre, workDate, odo)

	return nil
}
