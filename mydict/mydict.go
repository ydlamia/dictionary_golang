package mydict

import (
	"errors"
	"fmt"
)

// type Money int
// Money(1) = money=1

// Dictionary
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errExist = errors.New("It's already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err != nil{
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errExist
	}
	return nil
}

func Test1() {
	fmt.Println("Test1")
	Test2()
	defer fmt.Println("Test1 defer")
}

func Test2() {
	defer fmt.Println("Test2 defer")
	fmt.Println("Test2")

}
