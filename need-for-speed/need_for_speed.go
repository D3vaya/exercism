package speed

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func (car Car) PrintAttr() {
	fmt.Print("---ATTRIBUTES---\n")
	fmt.Printf("🔋 --> %v\n", car.battery)
	fmt.Printf("🪫  --> %v\n", car.batteryDrain)
	fmt.Printf("🏎  --> %v\n", car.speed)
	fmt.Printf("🛣  --> %v\n", car.distance)
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		speed:        speed,
		battery:      100,
		batteryDrain: batteryDrain,
	}
	return car
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	nTrak := Track{
		distance: distance,
	}
	return nTrak
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	fmt.Printf("🛞 %v veces\n", (car.battery / car.batteryDrain))
	fmt.Printf("Eso me da %v kilometros\n", (car.speed * (car.battery / car.batteryDrain)))
	fmt.Printf("🛣  %v kilometros\n", track.distance)

	if car.battery < car.batteryDrain {
		return false
	}

	distance := (car.speed * (car.battery / car.batteryDrain))

	if distance >= track.distance {
		return true
	}
	return false
}

func main() {
	// speed := 5
	// batteryDrain := 2
	// car := NewCar(speed, batteryDrain)
	// distance := 800
	// raceTrack := NewTrack(distance)
	// fmt.Println(car)
	// fmt.Println(raceTrack)
	// speed := 5
	// batteryDrain := 2
	// car := NewCar(speed, batteryDrain)
	// car = Drive(car)
	// fmt.Print(car)
	speed := 5
	batteryDrain := 7
	car := NewCar(speed, batteryDrain)
	Drive(car)
	//INPUT {battery:3 batteryDrain:7 speed:5 distance:0}
	//OUTPUT {battery:-4 batteryDrain:7 speed:5 distance:5};
	//SUCCESS {battery:3 batteryDrain:7 speed:5 distance:0}

	// distance := 100
	// raceTrack := NewTrack(distance)

	// isFinish := CanFinish(car, raceTrack)
	// fmt.Printf("Terminara la carrera ? %v\n", isFinish)
	// Output: true
}
