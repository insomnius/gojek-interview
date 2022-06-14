package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/insomnius/gojek-interview/entity"
)

type driver struct {
	driverPool      map[string]*entity.Driver
	availableDriver []string
}

func (d *driver) RegisterDriver(driverID string) error {
	if _, ok := d.driverPool[driverID]; ok {
		return errors.New("driver already registered")
	}

	d.driverPool[driverID] = &entity.Driver{
		CompleteBooking:          0,
		CompleteDistanceTraveled: 0,
		ID:                       driverID,
	}

	d.availableDriver = append(d.availableDriver, driverID)

	return nil
}

func (d *driver) AvailableDriver() []string {
	return d.availableDriver
}

func (d *driver) TakeDriver(driverID string) (*entity.Driver, error) {
	var driverNumber int

	for k, v := range d.availableDriver {
		if v == driverID {
			driverNumber = k
		}
	}

	if len(d.availableDriver) == 1 {
		d.availableDriver = []string{}
	} else {
		d.availableDriver = append(d.availableDriver[:driverNumber], d.availableDriver[driverNumber+1:]...)
	}

	return d.driverPool[driverID], nil
}

func (d *driver) ReturnDriver(driver *entity.Driver) {
	d.availableDriver = append(d.availableDriver, driver.ID)
}

func (d *driver) DriverList() map[string]*entity.Driver {
	return d.driverPool
}

func (d *driver) FindDriverWithDistanceTraveledGtThan(distance string) (*entity.Driver, error) {
	km, err := strconv.Atoi(distance)
	if err != nil {
		return nil, fmt.Errorf("Invalid KM arguments")
	}

	for _, driver := range d.driverPool {
		if driver.CompleteDistanceTraveled > km {
			return driver, nil
		}
	}

	return nil, errors.New("No driver founds with that criteria")
}
