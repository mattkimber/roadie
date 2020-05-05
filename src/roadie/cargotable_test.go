package roadie

import (
	"testing"
)

func TestCargoTable_Write(t *testing.T) {
	cargoTable := CargoTable{
		Cargo: []string{"ABCD", "DEFG"},
	}

	testTemplate(t, "testdata/output/cargo_table.nml", cargoTable)
}
