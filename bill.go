package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bills

func newBill(name string) bill {
	x := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0}

	return x
}

// add Item

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
	fmt.Printf("Item added is %v: $%0.2f \n", name, price)
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "Tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total+b.tip)

	return fs
}

// update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip
	fmt.Printf("The tip of %0.2f was added \n", b.tip)
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to the file")
}
