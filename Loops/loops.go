package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is ", x)
		x++
	}
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is ", i)
	}
	names := []string{"sun", "Shivam", "moon", "mountain", "river"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for index, value := range names { // if we don't want to use index replace it with _
		fmt.Printf("Value at index %v is %v\n", index, value)
	}
}
