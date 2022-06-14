package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/insomnius/gojek-interview/entity"
)

type booking struct {
	driver   Driver
	bookings map[string]*entity.Booking
}

func (b *booking) BookDriver(distance string) error {
	if len(b.driver.AvailableDriver()) == 0 {
		return fmt.Errorf("No driver available at the moment...")
	}

	km, err := strconv.Atoi(distance)
	if err != nil {
		return fmt.Errorf("Invalid KM arguments")
	}

	var driverNumber int

	if len(b.driver.AvailableDriver()) == 1 {
		driverNumber = 0
	} else {
		driverNumber = rand.Intn(len(b.driver.AvailableDriver()) - 1)
	}

	driver, _ := b.driver.TakeDriver(b.driver.AvailableDriver()[driverNumber])

	bookingID := fmt.Sprintf("booking-%d", len(b.bookings)+1)
	b.bookings[bookingID] = &entity.Booking{
		Distance: km,
		Complete: false,
		Driver:   driver,
	}

	return nil
}

func (b *booking) CompleteBooking(bookingID string) error {
	if booking, ok := b.bookings[bookingID]; ok {
		if booking.Complete {
			return fmt.Errorf("Booking already complete")
		}

		booking.Complete = true
		booking.Driver.CompleteBooking++
		booking.Driver.CompleteDistanceTraveled += booking.Distance
		b.driver.ReturnDriver(booking.Driver)
		return nil
	}

	return errors.New("booking is not valid")
}

func (d *booking) BookingList() map[string]*entity.Booking {
	return d.bookings
}
