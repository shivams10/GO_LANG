package main

import "fmt"

func main() {
	// Strings
	var nameOne string = "nameOne : it is of strict type."
	var nameTwo = "\nnameTwo: not defining the type but it implicitly define it as a string"
	var nameThree string
	nameFour := "\nnameFour : shorthand way to define a variable in go" //* this can't be done out of the main function scope
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// Integers
	var numOne int = 1
	var numTwo = 2
	var numThree int
	numFour := 4

	// Bits & memory
	var numFive int8 = 125 //* only allows number ranging inm 8 bit
	var numSix uint = 250 //*will not allow -ve numbers and we will get extra range of +ve numbers.
	
	// float
	var numSeven float32 = 10.20
	var numEight float32 = 100000000000.99 //* will be using float64 as it has higher precision

	fmt.Println(numOne, numTwo, numThree, numFour, numFive, numSix, numSeven, numEight)

}
