package mydict

import (
	"errors"
	"log"
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

func (d Dictionary) Update(word, def string) {
	_, err := d.Search(word)
	if err != nil {
		log.Fatalln(err)
	}
	d[word] = def
}

func (d Dictionary) Delete(word string) {
	_, err := d.Search(word)
	if err != nil {
		log.Fatalln(err)
	}
	delete(d, word)
}
