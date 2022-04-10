package interest

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance >= 0 {
		if balance < 1000 {
			return 0.5
		}
		if balance >= 1000 && balance < 5000 {
			return 1.621
		}
		if balance >= 5000 {
			return 2.475
		}
	}
	return 3.213
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interest := InterestRate(balance)
	percent := (interest * float32(balance)) / 100
	return float64(percent)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := InterestRate(balance)
	balanceUpdate := (balance * float64(interest)) / 100
	b := balance + balanceUpdate
	return b
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	//balance + interés por saldo inicial
	//saldo después de un año + interés por saldo después de un año
	years := 0
	for {
		if years > 0 {
			balance = AnnualBalanceUpdate(balance)
		}
		if balance >= targetBalance {
			return years
		}
		years++
	}
}

func main() {
	//fmt.Println(InterestRate(200.75))
	//fmt.Println(Interest(200.75))
	//fmt.Println(AnnualBalanceUpdate(200.75))
	balance := 200.75
	targetBalance := 214.88
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
	//YearsBeforeDesiredBalance(balance, targetBalance)
	// Output: 14
}
