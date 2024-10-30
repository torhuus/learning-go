package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	}

	if balance >= 0 && balance < 1000 {
		return 0.5
	}

	if balance >= 1000 && balance < 5000 {
		return 1.621
	}

	if balance >= 5000 {
		return 2.475
	}

	return 0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	if balance < 0 {
		return balance * (3.213 / 100)
	}

	if balance >= 0 && balance < 1000 {
		return balance * (0.5 / 100)
	}

	if balance >= 1000 && balance < 5000 {
		return balance * (1.621 / 100)
	}

	if balance >= 5000 {
		return balance * (2.475 / 100)
	}

	return balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	accBalance := balance

	yearsToTarget := 0

	for i := 0; i < 100; i++ {
		if accBalance < targetBalance {
			if targetBalance-accBalance < 0.00001 {
				break
			} else {
				accBalance = accBalance + Interest(accBalance)
				yearsToTarget = yearsToTarget + 1
			}
		} else {
			break
		}
	}
	return yearsToTarget
}
