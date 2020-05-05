package roadie

import (
	"assets"
	"bytes"
	"testing"
)

func testTemplate(t *testing.T, filename string, entity Entity) {
	expected, err := assets.GetInput(filename)
	if err != nil {
		t.Errorf("could not load test data: %v", err)
	}

	var output bytes.Buffer
	if err = entity.Write(&output); err != nil {
		t.Errorf("could not write template: %v", err)
	}

	if result := output.String(); result != expected {
		t.Errorf("template output - expected:\n%s\n\ngot:\n%s", expected, result)
	}
}
