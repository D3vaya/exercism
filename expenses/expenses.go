package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	s := []Record{}
	for _, v := range in {
		if predicate(v) {
			s = append(s, v)
		}
	}
	return s
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	counter := 0.0
	for _, v := range in {
		if v.Day >= p.From && v.Day <= p.To {
			counter += v.Amount
		}
	}
	return counter
}

var customError = errors.New("unknown category entertainment")

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	counter := 0.0
	inF := Filter(in, ByCategory(c))
	if len(inF) == 0 {
		return 0, customError
	}
	for _, v := range inF {
		if v.Day >= p.From && v.Day <= p.To {
			counter += v.Amount
		}

	}
	return counter, nil
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
	return r.Day == 1
}

func main() {
	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	num1, err1 := CategoryExpenses(records, p1, "entertainment")
	fmt.Println(num1, err1)
	// Output: 0, error(unknown category entertainment)

	num2, err2 := CategoryExpenses(records, p1, "rent")
	fmt.Println(num2, err2)
	// Output: 1300, nil

	num3, err3 := CategoryExpenses(records, p2, "rent")
	fmt.Println(num3, err3)

}
