package main

import (
	"fmt"
	"time"
)

func main() {
	customers := getData()
	fmt.Println(customers)
}

// A basic function that will retrieve data
func getData() (customers []string) {
	// Create an array of customer names
	customers = []string{"John Doe", "Jane Doe", "Robert Smith"}

	// Add elements to the slice
	customers = append(customers, "Jonas Bird")
	customers = append(customers, "Peter Murphy")
	customers = append(customers, "Robert Smith")
	customers = append(customers, "Daniel Ash")
	customers = append(customers, "Kirk Brandon")
	customers = append(customers, "Nik Wade")
	customers = append(customers, "Jaz Coleman")
	customers = append(customers, "Carl McCoy")
	customers = append(customers, "Rozz William")
	customers = append(customers, "Gavin Friday")

	// Experiment with an infinite loop
	for {
		fmt.Println("Infinite loop 1")
		time.Sleep(time.Second)
	}

	// return the array to the calling function
	return customers
}
