package mydict

import (
	"errors"
)

// type Money int
// Money(1) = money=1

// Dictionary
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
