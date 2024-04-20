package main

import "fmt"

func main() {
	age := 24
	name := "Shivam"
	score := 99.559432
	// Print
	fmt.Print("2nd line\n") //* Print method needs \n to add a new line
	fmt.Print("3rd line")

	// Println
	fmt.Println("1st line") //* Println method auto adds new line
	fmt.Println("my name is", name , "and my age is", age)

	// Printf (formatted string) %_ = format specifier
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("my name is %q and my age is %q \n", name, age) //* for strings
	fmt.Printf("age is of type %T \n", age) //* for types
	fmt.Printf("my score is %0.2f\n", score) //* the point before f will round off the number

	// Sprintf (Save formatted string)
	str := fmt.Sprintf("my score is %0.2f", score) //* the point before f will round off the number
	fmt.Println("Save string is",str)
}
