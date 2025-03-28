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
		err := dbFuncs.CreateTyreRepairEntry(&tyreRepair)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tyre repair not logged")
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
	tyreID := ReadString("Please enter tyre Pos: ")
	tyre, err := dbFuncs.ReadTyrePos(tyreID)
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
	tyre, err := dbFuncs.ReadTyreID(tyreID)
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
	tyre, err := dbFuncs.ReadTyreID(tyreID)
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

func RemoveTyreWork(tyreID string) error {
	tyre, err := dbFuncs.ReadTyreID(tyreID)
	if err != nil {
		return err
	}

	tyre.Location = ReadString("Please enter location: ")
	tyre.Position = "NULL"

	dbFuncs.CreateTyreEntry(&tyre)

	return nil
}

func Rotate() error {
	rotate := ReadInt("Please select from the below list:")
	for i := 0; i <= len(models.Rotate); i++ {
		fmt.Println(i, models.Rotate[i])
	}

	switch rotate {
	case 1:
		return RotateTyres()
	case 2:
		return RotateTyreOnRim()

	}

	return nil
}

func RotateTyres() error {
	fleetNum := ReadString("Please enter fleetnum of tyre 1: ")
	tyre1 := ReadString("Please enter position of tyre 1: ")
	tyre2 := ReadString("Please enter position of tyre 2: ")

	tyrepos1 := fleetNum + tyre1
	tyrepos2 := fleetNum + tyre2

	tyretype1, err := dbFuncs.ReadTyrePos(tyrepos1)
	if err != nil {
		return err
	}
	tyretype2, err := dbFuncs.ReadTyrePos(tyrepos2)
	if err != nil {
		return err
	}

	tyretype1.Position, tyretype2.Position = tyretype2.Position, tyretype1.Position

	dbFuncs.CreateTyreEntry(&tyretype1)

	time.Sleep(1 * time.Second)

	dbFuncs.CreateTyreEntry(&tyretype2)

	return nil
}

func RotateTyreOnRim() error {
	fleetNum := ReadString("Please enter fleetnum: ")
	tyrePos := ReadString("Please enter position: ")

	tyrePosition := fleetNum + tyrePos
	tyre, err := dbFuncs.ReadTyrePos(tyrePosition)
	if err != nil {
		return err
	}

	workDate := ReadDate("Please enter work date: (dd/mm/yyyy)")
	odo := ReadInt("Please enter odo: ")

	err = dbFuncs.CreateTyreRotateEntry(&tyre, workDate, odo)
	if err != nil {
		return err
	}

	fmt.Println("Tyre rotated")

	return nil
}
