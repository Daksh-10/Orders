package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose from option(a - Add Item, s - Save Item, t - Give Tip) ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Type Item name: ", reader)
		price, _ := getInput("Type Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be in number")
			promptOptions(b)
		}
		b.addItem(name, p)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be in number")
			promptOptions(b)
		}
		b.updateTip(t)
		promptOptions(b)
	case "s":
		fmt.Println("You chose to save bill")
		b.save()

	default:
		fmt.Println("You didn't chose from specified options")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
