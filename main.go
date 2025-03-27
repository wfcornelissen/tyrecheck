package main

import (
	_ "github.com/wfcornelissen/tyrecheck/internal/cmd"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
)

func main() {
	dbFuncs.CreateTyresTable()

}
