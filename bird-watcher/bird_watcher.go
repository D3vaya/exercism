package birdwatcher

import "fmt"

// TotalBirdCount devuelve el conteo total de aves sumando
// los recuentos del día individual.
func TotalBirdCount(birdsPerDay []int) int {
	var counter int
	for _, num := range birdsPerDay {
		counter += num
	}
	return counter
}

// BirdsInWeek devuelve el conteo total de aves sumando
// solo los artículos que pertenecen a la semana dada.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var counter int
	indexWeeks := map[int]int{
		1: 0,
		2: 7,
		3: 14,
		4: 21,
	}
	indexCutWeeks := map[int]int{
		1: 6,
		2: 13,
		3: 20,
		4: 27,
	}
	initWeek := indexWeeks[week]
	cutWeek := indexCutWeeks[week]

	for i, numBirds := range birdsPerDay[initWeek:(cutWeek + 1)] {
		if i <= indexCutWeeks[week] {
			counter += numBirds
		}
	}
	return counter
}

// FixBirdCountLog devuelve el conteo de aves después de corregir
// el pájaro cuenta por días alternos.
func FixBirdCountLog(birdsPerDay []int) []int {
	var counter int = 0
	var errorRange int = 1
	for i, num := range birdsPerDay {
		if counter%2 == 0 {
			birdsPerDay[i] = num + errorRange
		}
		counter++
	}
	return birdsPerDay
}

func main() {
	// birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	// fmt.Println(TotalBirdCount(birdsPerDay))
	//  32; want 14
	birdsPerDay := []int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0}
	fmt.Println(BirdsInWeek(birdsPerDay, 1))
	//birdsPerDay := []int{2, 5, 0, 7, 4, 1}
	//fmt.Println(FixBirdCountLog(birdsPerDay))
	// => [3 5 1 7 5 1]

}
