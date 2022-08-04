package main

import "fmt"

func main() {
	customer := getData(1)
	fmt.Println(customer)
}

// A basic function that will retrieve data
func getData(customerID int)(customer string) {
	// Example of control statement
	if customerID == 1 {
		return "Jonas Bird"
	} else if customerID == 2 {
		return "John Doe"
	} else {
		return ""
	}

	}
