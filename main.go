package main

import (
	"fmt"

	"github.com/ydlamia/dictionary_golang/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	// dictionary["hello"] = "bye"
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	// dictionary.add()
	// dictionary.delete()
}
