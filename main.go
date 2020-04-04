package main

import (
	"fmt"

	"github.com/ydlamia/dictionary_golang/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{"first": "First word"}
	// dictionary["hello"] = "bye"
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// a := dictionary.Add("first", "second Word")
	// if a != nil {
	// 	fmt.Println(a)
	// }
	// b := dictionary.Add("second", "second Word")
	// if b != nil {
	// 	fmt.Println(b)
	// }
	dictionary := mydict.Dictionary{}
	deleteWord := "bye"
	dictionary.Add("hello", "Hello")
	dictionary.Add("bye", "Bye")
	dictionary.Update("hello", "Hi")
	fmt.Println(dictionary)
	dictionary.Delete(deleteWord)
	fmt.Println(dictionary)
}
