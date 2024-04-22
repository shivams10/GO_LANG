package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           5.04,
		"pie":            7.99,
		"salad":          6.89,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping through map
	for k, v := range menu {
		fmt.Println(k, " ", v)
	}

	// ints as key type
	phoneBook := map[int]string{
		9808088 : "abc",
		9203088 : "def",
		1308088 : "ghi",
	}
	fmt.Println(phoneBook)
	phoneBook[7887888] = "jkl"
	fmt.Println(phoneBook)
}
