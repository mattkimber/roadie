package dto

import (
	"encoding/json"
	"github.com/mattkimber/roadie/src/roadie"
)

type CargoTableDTO struct {
	Cargo []string
}

func (d *CargoTableDTO) GetCargoTable() (c roadie.CargoTable) {
	c.Cargo = d.Cargo
	return
}

func (d *CargoTableDTO) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &d.Cargo); err != nil {
		return err
	}

	return nil
}
