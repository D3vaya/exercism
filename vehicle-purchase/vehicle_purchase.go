package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determina si se necesita una licencia para conducir un tipo de vehículo. Solo "automóvil" y "camión" requieren una licencia.
func NeedsLicense(kind string) bool {

	kinds := map[string]bool{
		"car":   true,
		"truck": true,
	}
	return kinds[kind]
}

// ChooseVehicle recomienda un vehículo para la selección. Siempre recomienda el vehículo que va primero en orden lexicográfico.
func ChooseVehicle(option1, option2 string) string {
	compare := strings.Compare(option2, option1)
	if compare == -1 {
		return option2 + " is clearly the better choice."
	} else {
		return option1 + " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// >3 -> 80%
	// 10 -> 50%
	// 3> && <= 10 -> 70%
	var calc float64
	if age < 3 {
		calc = float64(originalPrice*80) / 100
	}
	if age >= 10 {
		calc = float64(originalPrice*50) / 100
	}
	if age >= 3 && age < 10 {
		calc = float64(originalPrice*70) / 100
	}
	return calc
}

func main() {
	fmt.Println(NeedsLicense("truck"))
	fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	fmt.Println(CalculateResellPrice(1000, 1))
	// Output: 800
	fmt.Println(CalculateResellPrice(1000, 5))
	// Output: 700
	fmt.Println(CalculateResellPrice(1000, 15))
	// Output: 500
}
