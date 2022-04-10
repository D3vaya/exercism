package cards

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ns := []int{slice[1], slice[5], slice[8]}
	return ns
}

// GetItem recupera un elemento de un segmento en una posición dada.
// Si el índice está fuera de rango, queremos que devuelva -1.
func GetItem(slice []int, index int) int {
	fmt.Printf("Index -> %v, Len -> %v\n", index, len(slice))
	if index >= 0 && index <= len(slice)-1 {
		return slice[index]
	}
	return -1
}

// SetItem escribe un elemento en un slice en una posición determinada sobrescribiendo un valor existente.
// Si el índice está fuera de rango, se debe agregar el valor
func SetItem(slice []int, index, value int) []int {
	fmt.Printf("slice %v index %v value %v\n", len(slice)-1, index, value)
	if index < 0 || index > len(slice)-1 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	if len(value) > 0 {
		slice = append(value, slice...)
	}
	return slice
}

// RemoveItem elimina un elemento de un segmento modificando el segmento existente.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice)-1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// fmt.Println(FavoriteCards())

	fmt.Println(GetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, 7))
	// index := -1
	// newCard := 6
	// cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
	// fmt.Println(cards)
	// slice := []int{3, 2, 6, 4, 8}
	// cards := PrependItems(slice)
	// fmt.Println(cards)
	// cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	// fmt.Println(cards)
}
