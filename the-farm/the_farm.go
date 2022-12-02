package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if fodder < 0 && (err == ErrScaleMalfunction || err == nil) {
		negativeError := errors.New("negative fodder")
		return 0, negativeError
	}
	if err == ErrScaleMalfunction {
		return (fodder * 2) / float64(cows), nil
	} else if err != nil {
		someError := errors.New("non-scale error")
		return 0, someError
	}

	if cows == 0 {
		zeroError := errors.New("division by zero")
		return 0, zeroError
	} else if cows < 0 {
		sillyError := &SillyNephewError{cows: cows}
		return 0, sillyError
	}

	return fodder / float64(cows), nil
}
