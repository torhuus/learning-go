package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	Message      string
	NumberOfCows int
}

func (i *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", i.NumberOfCows, i.Message)
}

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, numberOfCows int) (float64, error) {
	totalFood, err := f.FodderAmount(numberOfCows)
	if err != nil {
		return 0.0, err
	}

	fatFactor, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	foodPerCow := (totalFood * fatFactor) / float64(numberOfCows)

	return foodPerCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(f, numberOfCows)
	} else {
		return 0.0, errors.New("invalid number of cows")
	}
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{
			NumberOfCows: numberOfCows,
			Message:      "there are no negative cows",
		}
	}
	if numberOfCows == 0 {
		return &InvalidCowsError{
			NumberOfCows: numberOfCows,
			Message:      "no cows don't need food",
		}
	}
	return nil
}
