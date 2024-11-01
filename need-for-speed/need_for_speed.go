package speed

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.battery = car.battery - car.batteryDrain
	car.distance = car.distance + car.speed

	return car
}

type Shitman struct {
	shii int
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	newShitman := Shitman{shii: 9}
	// How many times can we run the car? rem battery div by drain
	r := car.battery / car.batteryDrain

	// Using the r, remaining full drives times the speed gives us the total distance possible to drive
	r = r * car.speed

	// r >= track.distance checks if the drive distance remaining is less than track distance and returns true iif so
	return r >= track.distance
}
