package gross

var (
	units = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
)

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	v, exists := units[unit]
	if !exists {
		return false
	}
	_, existsBill := bill[item]
	if existsBill {
		bill[item] += v
	} else {
		bill[item] = v
	}

	return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Devuelve falso si el artículo dado no está en la factura
	_, exists := bill[item]
	if !exists {
		return false
	}
	// Devuelve falso si la unidad dada no está en el mapa de unidades.
	v2, existsUnits := units[unit]
	if !existsUnits {
		return false
	}
	// Devuelve false si la nueva cantidad sería menor que 0.
	isMinor := bill[item] - v2
	//fmt.Println("LOG", bill[item], v2)
	if isMinor < 0 {
		return false
	}
	bill[item] -= v2
	// Si la nueva cantidad es 0, elimine completamente el artículo de la factura y luego devuelva verdadero.
	// De lo contrario, reduzca la cantidad del artículo y devuelva verdadero.
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, exists := bill[item]
	if !exists {
		return 0, false
	}
	return bill[item], true

}

func main() {
	// units := Units()
	// fmt.Println(units)
	// bill := NewBill()
	// units := Units()
	// ok := AddItem(bill, units, "carrot", "dozen")
	// fmt.Println(ok)

	// bill := NewBill()
	// units := Units()
	// ok := RemoveItem(bill, units, "peas", "half_of_a_dozen")
	// fmt.Println(ok)
	// Output: false (because there are no carrots in the bill)

}
