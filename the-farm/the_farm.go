package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	numbersOfCow int
}

func (s *SillyNephewError) Error() error {
	msm := fmt.Sprintf("silly nephew, there cannot be %v cows", s.numbersOfCow)
	return errors.New(msm)
}

var NegativeFodder = errors.New("negative fodder")
var DivisioByZero = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	//1. Obtenga la cantidad de forraje del FodderAmountmétodo .
	if err != nil {
		if err.Error() == "non-scale error" {
			return 0.0, err
		}
		if err.Error() == "sensor error" && fodder > 0 {
			fodder = fodder * 2.0
		}
	}

	//2. Devolver un error por forraje negativo
	if fodder < 0.0 {
		return 0.0, NegativeFodder
	}

	//3. Evitar la división por cero
	if cows == 0.0 {
		return 0.0, DivisioByZero
	}

	//4. Manejar vacas negativas
	s := SillyNephewError{
		numbersOfCow: cows,
	}
	if cows < 0.0 {
		return 0.0, s.Error()
	}

	fodderWithCows := fodder / float64(cows)

	return fodderWithCows, nil
}

func main() {
	//twentyFodderNoError := {}
	// fodder, _ := DivideFood(twentyFodderNoError, 10)
	// fmt.Println(fodder)
}
