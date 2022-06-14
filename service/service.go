package service

import "github.com/insomnius/gojek-interview/entity"

type Driver interface {
	RegisterDriver(driverID string) error
	AvailableDriver() []string
	TakeDriver(driverID string) (*entity.Driver, error)
	ReturnDriver(driver *entity.Driver)
	DriverList() map[string]*entity.Driver
	FindDriverWithDistanceTraveledGtThan(distance string) (*entity.Driver, error)
}

func NewDriver(driverPool map[string]*entity.Driver) Driver {
	return &driver{
		driverPool:      driverPool,
		availableDriver: []string{},
	}
}

type Booking interface {
	BookDriver(distance string) error
	CompleteBooking(bookingID string) error
	BookingList() map[string]*entity.Booking
}

func NewBooking(bookings map[string]*entity.Booking, driver Driver) Booking {
	return &booking{
		bookings: bookings,
		driver:   driver,
	}
}
