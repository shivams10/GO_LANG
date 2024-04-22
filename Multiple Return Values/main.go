package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	str := strings.ToUpper(n)
	names := strings.Split(str , " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1{
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main()  {
	firstName , secondName := 	getInitials("Shivam Shukla")
	fmt.Println(firstName, secondName)
}