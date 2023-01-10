package internal

import (
	"matcher/pkg"
)

type Coordinates struct {
	Latitude  float64 `query:"lat"`
	Longitude float64 `query:"long"`
}

func (c *Coordinates) Validate() error {
	if c.Latitude < -90 || c.Latitude > 90 || c.Longitude < -180 || c.Longitude > 180 {
		return pkg.ErrInvalidRequest
	}

	return nil
}
