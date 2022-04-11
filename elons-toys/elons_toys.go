package elon

import "fmt"

// TODO: define the 'Drive()' method

func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	fmt.Println(fmt.Sprintf("Driven %v meters", c.distance))
	return fmt.Sprintf("Driven %v meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%s", c.battery, "%")
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {

	fmt.Printf("🛞 %v veces\n", (c.battery / c.batteryDrain))
	fmt.Printf("Eso me da %v kilometros\n", (c.speed * (c.battery / c.batteryDrain)))
	fmt.Printf("🛣  %v kilometros\n", trackDistance)

	if c.battery < c.batteryDrain {
		return false
	}

	distance := (c.speed * (c.battery / c.batteryDrain))

	if distance >= trackDistance {
		return true
	}
	return false

}
