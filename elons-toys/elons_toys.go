package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery-c.batteryDrain < 0 {
		return
	}

	c.battery = c.battery - c.batteryDrain
	c.distance = c.distance + c.speed
	return
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	remainingRange := (c.battery / c.batteryDrain) * c.speed

	if trackDistance <= remainingRange {
		return true
	}

	return false
}
