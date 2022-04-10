package cars

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * float64(successRate) / 100
}

// CalculateWorkingCarsPerMinute calcula cuántos coches en funcionamiento hay
// producido por la línea de montaje cada minuto
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	percent := (float64(productionRate) * successRate) / 100
	return int(percent / 60)
}

// CalculateCost calcula el costo de producir la cantidad dada de autos
func CalculateCost(carsCount int) uint {
	const OptimalNumber = 10
	diff := carsCount / OptimalNumber
	mod := carsCount % OptimalNumber
	return uint((diff * 95000) + (mod * 10000))

}

func main() {
	rate := CalculateWorkingCarsPerHour(1547, 90)
	fmt.Println(rate)
	// Output: 1392.3

	rate2 := CalculateWorkingCarsPerMinute(1105, 90)
	fmt.Println(rate2)
	// Output: 16
	cost := CalculateCost(37)
	fmt.Println(cost)
	// Output: 355000
	cost = CalculateCost(21)
	fmt.Println(cost)
	// Output: 200000

}
