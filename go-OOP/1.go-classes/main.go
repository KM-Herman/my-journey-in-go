package main

import "fmt"

type Company struct {
	Id      string
	Name    string
	Country string
}

func NewCompany(id string, name string, country string) Company {
	return Company{
		Id:      id,
		Name:    name,
		Country: country,
	}
}

func main() {

	var company Company = NewCompany("234", "Kikuu", "China")

	fmt.Println(company)
}
