package geo

import "errors"

type Coordinates struct {
	lat float64
	lon float64
}
type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}

func (l *Landmark) SetLat(latitude float64) {
	l.Coordinates.lat = latitude
}

func (l *Landmark) SetLon(longitude float64) {
	l.Coordinates.lon = longitude
}
