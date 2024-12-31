package gameError

import "errors"

var NoBodyFound = errors.New("Could not find body of snake")

func CouldNotEat(err error) error {
	return errors.New("Could")
}
