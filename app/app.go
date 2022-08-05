package main

import "fmt"

func main() {
	customers := getData()
	fmt.Println(customers)
}

// A basic function that will retrieve data
func getData()(customers [2]string) {
	// Create an array of customer names
	customer := "Jonas Bird"

	// Add the entry to an array
	customers[0] = customer

	// Assign a value directly to the array
	customers[1] = "John Doe"

	// return the array to the calling function
	return customers
	}

