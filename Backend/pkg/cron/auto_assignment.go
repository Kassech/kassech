package cron

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"kassech/backend/pkg/config"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

var (
	mutex     sync.Mutex
	threshold = 4
	batchSize = 12 // Assign in bulk of 12 passengers per car
)

type IgnoredPath struct {
	PathID     uint
	Passengers []models.Passenger
}

func RunAssignmentLoop() {
	passengerRepo := repository.PassengerRepository{}
	pathRepo := repository.PathRepository{}
	carRepo := repository.VehicleRepository{} // Assuming there's a car repository
	for {
		log.Println("Starting assignment loop...")
		assignDriversToQueues(passengerRepo, pathRepo, carRepo)

		log.Println("Waiting 2 seconds before next run...")
		time.Sleep(5 * time.Second)
	}
}
func assignDriversToQueues(passengerR repository.PassengerRepository, pathR repository.PathRepository, carR repository.VehicleRepository) {
	mutex.Lock()
	log.Println("Acquired lock for driver assignment...")
	defer func() {
		mutex.Unlock()
		log.Println("Released lock after driver assignment...")
	}()

	log.Println("Fetching passengers...")
	passengers, total, err := passengerR.GetAll(1, 100, "", 0)
	if err != nil {
		log.Printf("Error fetching passengers: %v", err)
		return
	}
	log.Printf("Fetched %d passengers, total: %d", len(passengers), total)

	pathPassengers := make(map[uint][]models.Passenger)
	for _, p := range passengers {
		pathPassengers[p.PathID] = append(pathPassengers[p.PathID], p)
	}

	var ignoredPaths []IgnoredPath
	for pathID, pList := range pathPassengers {

		log.Printf("Processing path %d with %d passengers", pathID, len(pList))
		if len(pList) < threshold {
			ignoredPaths = append(ignoredPaths, IgnoredPath{PathID: pathID, Passengers: pList})
			log.Printf("Ignoring path %d with %d passengers", pathID, len(pList))
			continue
		}
		path, err := pathR.FindByID(pathID)
		if err != nil {
			log.Printf("Error fetching path %d: %v", pathID, err)
			continue
		}

		log.Printf("Assigning passengers in bulk for path")
		assignPassengersInBulk(path, pList, carR, passengerR)
	}

	log.Println("Assignment cycle completed.")
}
func assignPassengersInBulk(path *models.Path, passengers []models.Passenger, carR repository.VehicleRepository, passengerR repository.PassengerRepository) {
	log.Printf("Fetching available cars for path %d...", path.ID)
	availableCars, err := carR.GetNearestAvailableCars(path.Route.StationA.ID, 5*60*1000)
	if err != nil {
		log.Printf("Error fetching available cars for path %d: %v", path.ID, err)
		return
	}

	log.Printf("Found %d available cars for path %d", len(availableCars), path.ID)
	if len(availableCars) == 0 {
		log.Printf("No available cars for path %d", path.ID)
		return
	}

	// Log car details for debugging
	log.Printf("Available cars for path %d:", path.ID)
	for _, car := range availableCars {
		log.Printf("Car ID: %d, Type: %+v, Capacity: %d", car.ID, car.Type, car.Type.Capacity)
	}

	if len(passengers) == 0 {
		log.Printf("No passengers to assign for path %d", path.ID)
		return
	}

	for _, car := range availableCars {
		// Skip cars with invalid capacity
		if car.Type.Capacity <= 0 {
			log.Printf("Skipping car %d: invalid capacity %d", car.ID, car.Type.Capacity)
			continue
		}

		if len(passengers) == 0 {
			log.Printf("No more passengers left to assign for path %d", path.ID)
			break
		}

		assignable := min(len(passengers), min(batchSize, int(car.Type.Capacity)))
		log.Printf("Assignable passengers: %d (batchSize: %d, car capacity: %d)", assignable, batchSize, car.Type.Capacity)

		if assignable <= 0 {
			log.Printf("No passengers to assign to car %d (assignable: %d)", car.ID, assignable)
			continue
		}

		selectedPassengers := passengers[:assignable]
		passengers = passengers[assignable:]

		log.Printf("Assigning %d passengers to Car %d on path %d", assignable, car.ID, path.ID)
		err := passengerR.AssignToCarAndRemove(selectedPassengers, car, path.ID)
		if err != nil {
			log.Printf("Failed to assign passengers to car %d: %v", car.ID, err)
		} else {
			pathBytes, err := json.Marshal(path)
			if err != nil {
				log.Printf("Error marshalling message for car %d: %v", car.ID, err)
				continue
			}

			config.ConnManager.SendToUser(car.Driver.User.ID, pathBytes)

			log.Printf("Successfully assigned %d passengers to Car %d on path %d", assignable, car.ID, path.ID)
		}
	}

	log.Printf("Finished assigning passengers for path %d", path.ID)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
