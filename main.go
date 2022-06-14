package main

import (
	"fmt"

	"github.com/insomnius/gojek-interview/entity"
	"github.com/insomnius/gojek-interview/service"
)

var driverPool = map[string]*entity.Driver{}
var availableDriver = []string{}
var bookings = map[string]*entity.Booking{}

func main() {
	fmt.Println("Gojek Driver CLI")

	driverService := service.NewDriver(driverPool)
	bookingService := service.NewBooking(bookings, driverService)

	for {
		var command string
		var args string

		_, err := fmt.Scanln(&command, &args)
		if err != nil && err.Error() != "unexpected newline" {
			panic(err)
		}

		switch command {
		case "register_driver":
			if err := driverService.RegisterDriver(args); err != nil {
				fmt.Println("Error:", err)
			}
		case "driver_list":
			for id, driver := range driverService.DriverList() {
				fmt.Println("ID:", id, "Driver KM:", driver.CompleteDistanceTraveled, "Driver Complete Booking:", driver.CompleteBooking)
			}
		case "dispatch_driver_for_a_booking":
			if err := bookingService.BookDriver(args); err != nil {
				fmt.Println("Error:", err)
			}
		case "booking_completed_distance_gt_10":
			if driver, err := driverService.FindDriverWithDistanceTraveledGtThan("10"); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("ID:", driver.ID, "Driver KM:", driver.CompleteDistanceTraveled, "Driver Complete Booking:", driver.CompleteBooking)
			}
		case "booking_list":
			for id, booking := range bookingService.BookingList() {
				fmt.Println("ID:", id, "Driver ID:", booking.Driver.ID, "Booking:", booking)
			}
		case "complete_booking":
			if err := bookingService.CompleteBooking(args); err != nil {
				fmt.Println("Error:", err)
			}
		case "exit":
			fmt.Println("Exiting application")
			return
		default:
			fmt.Println("Command is not available")
		}
	}
}
