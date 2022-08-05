package main

import "fmt"

func main() {
	customers := getData()
	fmt.Println(customers)
}

// A basic function that will retrieve data
func getData()(customers []string) {
	// Create an array of customer names
	customers = []string{"John Doe", "Jane Doe", "Robert Smith"}

	// Add elements to the slice
	customers = append(customers, "Jonas Bird")

	// return the array to the calling function
	return customers
	}

