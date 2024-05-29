package main

import "fmt"

func updateName(nameParams *string) {
	*nameParams = "SHIVAM"
}

func main() {
	name := "shivam"
	namePointer := &name
	updateName(namePointer)
	fmt.Println(name)
}
