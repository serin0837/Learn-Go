package main

import (
"fmt"
"github.com/serin0837/learngo/mydict"
)


func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	fmt.Println(dictionary)
	dictionary.Delete(baseWord)
	fmt.Println(dictionary)
}
