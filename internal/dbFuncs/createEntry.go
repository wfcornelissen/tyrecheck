package dbFuncs

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func CreateTyreEntry(tyre *models.Tyre) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tyres (id, size, brand, model, supplier, price, position, location, state, condition, startingTread, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		tyre.ID,
		tyre.Size,
		tyre.Brand,
		tyre.Model,
		tyre.Supplier,
		tyre.Price,
		tyre.Position,
		tyre.Location,
		tyre.State,
		tyre.Condition,
		tyre.StartingTread,
		tyre.Archived)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: tyres.id" {
			fmt.Println("--------------------------------")
			fmt.Println("Tyre with this id already exists")
			fmt.Println("--------------------------------")
			return errors.New("tyreid already exists")
		} else {
			fmt.Println("Error creating tyre entry: ")
			return err
		}
	}

	fmt.Println("Tyre entry created")

	return nil
}

func CreateTruckEntry(truck *models.Truck) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO trucks (fleetNum, vin, reg, make, model, year, odo, scrap, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		truck.FleetNum,
		truck.VIN,
		truck.Reg,
		truck.Make,
		truck.Model,
		truck.Year,
		truck.Odo,
		truck.Scrap,
		truck.Archived)
	if err != nil {
		fmt.Println("Error creating truck entry")
		return err
	}

	fmt.Println("Truck entry created")
	return nil
}

func CreateTrailerEntry(trailer *models.Trailer) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO trailers (fleetNum, vin, reg, make, model, year, scrap, archived) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		trailer.FleetNum,
		trailer.VIN,
		trailer.Reg,
		trailer.Make,
		trailer.Model,
		trailer.Year,
		trailer.Scrap,
		trailer.Archived)
	if err != nil {
		fmt.Println("Error creating trailer entry")
		return err
	}

	fmt.Println("Trailer entry created")

	return nil
}

func CreateCombinationEntry(combination *models.Combination) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO combinations (truckFleetNum, trailerFleetNum) VALUES (?, ?)",
		combination.TruckFleetNum,
		combination.TrailerFleetNum)
	if err != nil {
		fmt.Println("Error creating combination entry")
		return err
	}

	fmt.Println("Combination entry created")

	return nil
}

func CreateTyreCheckEntry(tyreCheck *models.TyreWork) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tyrechecks (id, tyreID, workDate, position, odo) VALUES (?, ?, ?, ?, ?)",
		tyreCheck.ID,
		tyreCheck.TyreID,
		tyreCheck.WorkDate,
		tyreCheck.Position,
		tyreCheck.Odo)

	if err != nil {
		fmt.Println("Error creating tyre check entry")
		return err
	}

	fmt.Println("Tyre check entry created")

	return nil
}

func CreateTyreRepairEntry(tyreRepair *models.TyreWork) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tyrerepairs (id, tyreID, workDate, position, odo) VALUES (?, ?, ?, ?, ?)",
		tyreRepair.ID,
		tyreRepair.TyreID,
		tyreRepair.WorkDate,
		tyreRepair.Position,
		tyreRepair.Odo)

	if err != nil {
		fmt.Println("Error creating tyre repair entry")
		return err
	}

	fmt.Println("Tyre repair entry created")
	return nil
}

func CreateRetreadSentEntry(retreadSent *models.Tyre, WorkDate time.Time, Odo int) error {
	// Will need logic in workFuncs and readEntry.go
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO retreadSent (id, tyreID, workDate, position, odo) VALUES (?, ?, ?, ?, ?)",
		retreadSent.ID,
		WorkDate,
		Odo)

	if err != nil {
		fmt.Println("Error creating retread sent entry")
		return err
	}

	fmt.Println("Retread sent entry created")

	return nil
}

func CreateRetreadReceivedEntry(retreadReceived *models.Tyre, WorkDate time.Time, Odo int) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO retreadReceived (id, tyreID, dateReceived, odo) VALUES (?, ?, ?, ?)",
		retreadReceived.ID,
		WorkDate,
		Odo)

	if err != nil {
		fmt.Println("Error creating retread received entry")
		return err
	}

	fmt.Println("Retread received entry created")

	return nil
}

func CreateRetreadScrapEntry(retreadScrap *models.Tyre, WorkDate time.Time, Odo int) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO retreadScrap (id, tyreID, dateScraped, odo) VALUES (?, ?, ?, ?)",
		retreadScrap.ID,
		WorkDate,
		Odo)

	if err != nil {
		fmt.Println("Error creating retread scrap entry")
		return err
	}

	fmt.Println("Retread scrap entry created")

	return nil
}

func CreateTyreRotateEntry(tyreRotate *models.Tyre, WorkDate time.Time, Odo int) error {
	db, err := sql.Open("sqlite3", "./tyrecheck.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tyrerotates (id, tyreID, position, dateRotated, odo) VALUES (?, ?, ?, ?, ?)",
		tyreRotate.ID,
		tyreRotate.Position,
		WorkDate,
		Odo)

	if err != nil {
		fmt.Println("Error creating tyre rotate entry")
		return err
	}

	return nil
}
