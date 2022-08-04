package main

import "fmt"

func main() {
	customer := getData(1)
	fmt.Println(customer)
}

// A basic function that will retrieve data
func getData(customerID int)(customer string) {
	var firstName = "Jonas"
	lastName := "Bird"
	var fullName string
	fullName = firstName + " " + lastName
	return fullName
}
