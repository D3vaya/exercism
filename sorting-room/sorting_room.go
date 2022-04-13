package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %v.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	num, err := strconv.Atoi(fnb.Value())
	fmt.Println(num, err)
	if err != nil {
		return 0
	}
	if num < 10 {
		return 0
	}
	return num
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num, err := strconv.Atoi(fnb.Value())
	if err != nil {
		return fmt.Sprintf("This is a fancy box containing the number 0.0")
	}
	if num < 10 {
		return fmt.Sprintf("This is a fancy box containing the number 0.0")
	}

	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	n, ok := i.(int)
	if ok {
		return DescribeNumber(float64(n))
	}
	n64, ok := i.(float64)
	if ok {
		return DescribeNumber(n64)
	}
	nb, ok := i.(NumberBox)
	if ok {
		return DescribeNumberBox(nb)
	}
	fn, ok := i.(FancyNumberBox)
	if ok {
		return DescribeFancyNumberBox(fn)
	}

	return "Return to sender"

}
