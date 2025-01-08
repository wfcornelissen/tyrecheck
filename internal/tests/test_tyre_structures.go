package tests

import (
	"testing"

	"github.com/wfcornelissen/tyrecheck/internal/models"
)

func TestIsValidTyreSize(t *testing.T) {
	newTyre := models.Tyre{}
	newTyre.TyreSize = 315
	if !models.IsValidTyreSize(models.TyreSize(newTyre.TyreSize)) {
		t.Errorf("Tyre size is not valid")
	}
	newTyre.TyreSize = 385
	if !models.IsValidTyreSize(models.TyreSize(newTyre.TyreSize)) {
		t.Errorf("Tyre size is not valid")
	}
	newTyre.TyreSize = 375
	if models.IsValidTyreSize(models.TyreSize(newTyre.TyreSize)) {
		t.Errorf("Tyre size is valid")
	}
}
