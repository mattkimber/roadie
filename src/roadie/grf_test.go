package roadie

import (
	"testing"
)

func TestGrf_Write(t *testing.T) {
	grf := Grf{
		Identifier:           "ABCD",
		Version:              2,
		MinCompatibleVersion: 1,
	}

	testTemplate(t, "testdata/output/grf.nml", grf)
}
