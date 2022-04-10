package lasagna

import "fmt"

// TODO: define the 'PreparationTime()' function
// NOTE OK
func PreparationTime(layers []string, time int) int {
	var defaultTime int = 2
	if time == 0 {
		time = defaultTime
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {

	gramsNoodles := 50
	litersWater := 0.2

	gramsForLayers := 0
	litersforSauce := 0.0
	for _, v := range layers {
		if v == "noodles" {
			gramsForLayers += gramsNoodles
		}
		if v == "sauce" {
			litersforSauce += litersWater
		}
	}
	return gramsForLayers, litersforSauce
}

// TODO: define the 'AddSecretIngredient()' function
// NOTE OK
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

	//quantities 	-> {0.5, 250, 150, 3, 0.5}
	//portions 		-> 1
	//expect 			-> {0.25, 125, 75, 1.5, 0.25}
	var defaultPortions int = 2
	recipe := []float64{}
	for _, quantity := range quantities {

		num := quantity / float64(defaultPortions)
		recipe = append(recipe, num*float64(portions))
	}

	return recipe
}

func main() {
	//layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	// fmt.Println(PreparationTime(layers, 3))
	// fmt.Println(PreparationTime(layers, 0))
	//g, l := Quantities(layers)

	// friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	// myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	// AddSecretIngredient(friendsList, myList)

	// quantities := []float64{1.2, 3.6, 10.5}
	// scaledQuantities := ScaleRecipe(quantities, 4)
	// fmt.Println(scaledQuantities)
	// => []float64{ 2.4, 7.2, 21 }
	quantities2 := []float64{0.5, 250, 150, 3, 0.5}
	scaledQuantities2 := ScaleRecipe(quantities2, 1)
	fmt.Println(scaledQuantities2)
	// => []float64{0.25, 125, 75, 1.5, 0.25}
}
