package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jsgoecke/tesla"
)

func main() {
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     os.Getenv("TESLA_CLIENT_ID"),
			ClientSecret: os.Getenv("TESLA_CLIENT_SECRET"),
			Email:        os.Getenv("TESLA_USERNAME"),
			Password:     os.Getenv("TESLA_PASSWORD"),
		})
	if err != nil {
		panic(err)
	}

	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}
	vehicle := vehicles[0]
	status, err := vehicle.MobileEnabled()

	if err != nil {
		panic(err)
	}

	fmt.Println(status)
	// fmt.Println(vehicle.ChargeState())
	// fmt.Println(vehicle.ClimateState())
	// fmt.Println(vehicle.DriveState())
	// fmt.Println(vehicle.GuiSettings())
	// fmt.Println(vehicle.VehicleState())
	// fmt.Println(vehicle.HonkHorn())
	// fmt.Println(vehicle.FlashLights())
	// fmt.Println(vehicle.Wakeup())
	// fmt.Println(vehicle.OpenChargePort())
	// fmt.Println(vehicle.ResetValetPIN())
	// fmt.Println(vehicle.SetChargeLimitStandard())
	// fmt.Println(vehicle.SetChargeLimit(50))
	// fmt.Println(vehicle.StartCharging())
	// fmt.Println(vehicle.StopCharging())
	// fmt.Println(vehicle.SetChargeLimitMax())
	// fmt.Println(vehicle.StartAirConditioning())
	// fmt.Println(vehicle.StopAirConditioning())
	// fmt.Println(vehicle.UnlockDoors())
	// fmt.Println(vehicle.LockDoors())
	// fmt.Println(vehicle.SetTemprature(72.0, 72.0))
	// fmt.Println(vehicle.Start(os.Getenv("TESLA_PASSWORD")))
	// fmt.Println(vehicle.OpenTrunk("rear"))
	// fmt.Println(vehicle.OpenTrunk("front"))
	// fmt.Println(vehicle.MovePanoRoof("vent", 0))
	// fmt.Println(vehicle.MovePanoRoof("open", 0))
	// fmt.Println(vehicle.MovePanoRoof("move", 50))
	// fmt.Println(vehicle.MovePanoRoof("close", 0))
	// fmt.Println(vehicle.TriggerHomelink())

	// Take care with these, as the car will move
	// fmt.Println(vehicle.AutoparkForward())
	// fmt.Println(vehicle.AutoparkReverse())
	// Take care with these, as the car will move

	// Stream vehicle events
	eventChan, errChan, err := vehicle.Stream()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		for {
			select {
			case event := <-eventChan:
				eventJSON, _ := json.Marshal(event)
				fmt.Println(string(eventJSON))
			case err = <-errChan:
				fmt.Println(err)
				if err.Error() == "HTTP stream closed" {
					fmt.Println("Reconnecting!")
					eventChan, errChan, err = vehicle.Stream()
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
}
