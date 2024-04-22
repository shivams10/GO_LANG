package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string){
	fmt.Printf("Good bye %v \n", n)
}

func cycleNames (n[]string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea (r float64) float64{
	return math.Pi * r *r
}

func main() {
	// sayGreeting("Shivam")
	// sayGreeting("India")
	// sayBye("Bro")
	cycleNames([]string{"cloud", "sun", "moon"}, sayGreeting)
	cycleNames([]string{"cloud", "sun", "moon"}, sayBye)
	a1 := circleArea(10.5)
	a2 := circleArea(12.5)
	fmt.Printf("area of circle 1 is %0.3f & that of circle 2 is %0.3f", a1, a2)
}
