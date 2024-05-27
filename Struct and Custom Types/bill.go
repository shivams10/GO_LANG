package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// * function to make new  bills
func newBill(name string) bill {
	aBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return aBill
}

// Receiver function
