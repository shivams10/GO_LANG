package main

import "fmt"
func main()  {
	// var arg[3]int = [3]int{20, 25, 30}
	var arg = [3]int{20, 25, 30}
	fmt.Println(arg, len(arg))

	names:= []string{"angel", "Hong Kong", "good food"}
	names[2] = "healthy food"
	fmt.Println(names, len(names))	
	names = append(names, "India")
	fmt.Println(names, len(names))

	// Slices 
	newName := names[0:2]
	fmt.Println(newName, len(newName))

}