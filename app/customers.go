package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {
	jonas := Customer{FirstName: "Jonas", LastName: "Bird"}

	customers = append(customers, jonas,
		Customer{FirstName: "Rikk", LastName: "Agnew"},
		Customer{FirstName: "Peter", LastName: "Murphy"},
		Customer{FirstName: "Robert", LastName: "Smith"},
		Customer{FirstName: "Daniel", LastName: "Ash"},
		Customer{FirstName: "Kirk", LastName: "Brandon"},
		Customer{FirstName: "Nik", LastName: "Wade"},
		Customer{FirstName: "Jaz", LastName: "Coleman"},
		Customer{FirstName: "Carl", LastName: "McCoy"},
		Customer{FirstName: "Rozz", LastName: "Williams"},
		Customer{FirstName: "Gavin", LastName: "Friday"},
		Customer{FirstName: "Genesis", LastName: "Orrige"},
		Customer{FirstName: "Nivek", LastName: "Ogre"},
		Customer{FirstName: "All", LastName: "Jurgensen"},
	)
	return customers
}
