// package chance
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime siembra matemáticas/rand con la hora actual de la computadora
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie devuelve un int aleatorio d con 1 <= d <= 20
func RollADie() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	a := r.Intn(20 - 1)
	b := 20
	return a + rand.Intn(b-a+1)
}

// GenerateWandEnergy devuelve un float64 f aleatorio con 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	source := rand.NewSource(time.Now().UnixNano())
	ra := rand.New(source)
	fmt.Println(ra.Float64())
	var min float64 = 0
	var max float64 = 12.0
	r := min + rand.Float64()*(max-min)
	return r
}

// ShuffleAnimals devuelve un segmento con las ocho cadenas de animales en orden aleatorio
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}

func main() {
	//SeedWithTime()
	//d := RollADie()
	f := GenerateWandEnergy()
	fmt.Println(f)
}
