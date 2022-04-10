// Package weather packete condicion del tiempo.
package weather

// CurrentCondition condicion actual del tiempo.
var CurrentCondition string

// CurrentLocation  variable de la ubicacion actual.
var CurrentLocation string

// Forecast funcion que retorna la condicion del tiempo.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
