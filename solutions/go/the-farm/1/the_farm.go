package thefarm

import (
    "fmt"
    "errors"
)

type InvalidCowsError struct {
    numberOfCows int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

func DivideFood(fc FodderCalculator, nc int) (float64, error) {
    fodder, err := fc.FodderAmount(nc)
    if err != nil {
        return 0, err
    }
    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    fodderPerCow := (fodder / float64(nc)) * factor
    return fodderPerCow, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {
    if nc <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, nc)
}

func ValidateNumberOfCows(nc int) error {
    if nc < 0 {
        return &InvalidCowsError{
            numberOfCows: nc,
            message: "there are no negative cows",
        }
    }
    if nc == 0 {
        return &InvalidCowsError{
            numberOfCows: nc,
            message: "no cows don't need food",
        }
    }
    return nil
}
