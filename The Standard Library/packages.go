package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings
	greeting := " Hello there friends!"
	isContainHello := strings.Contains(greeting, "Hello")
	replaceWithHi := strings.ReplaceAll(greeting, "Hello", "Hi")
	upperCase := strings.ToUpper(greeting)
	splitString := strings.Split(greeting, " ")

	// sort
	scores := []int{100, 90, 89, 99, 88}
	sort.Ints(scores)
	scoreIndex := sort.SearchInts(scores, 90)
	names := []string{"sun", "Shivam", "moon", "mountain", "river"}
	sort.Strings(names)
	searchString := sort.SearchStrings(names, "Shivam")

	fmt.Println(greeting, isContainHello, replaceWithHi, upperCase, splitString)
	fmt.Println(scores, scoreIndex, names, searchString)
}
