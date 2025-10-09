package main

import "fmt"

type Company struct {
	Id string
	Name string
	Country string
}

// constructor
func NewCompany(id string, name string, country string) Company {
	if id == "" {
		panic("company ID cannot be empty")
	}
	if name == "" {
		panic("company name cannot be empty")
	}

	return Company{
		Id: id,
		Name: name,
		Country: country,
	}
}

// Getter methods which provides controlled access to private fields
func (c Company) GetID() string {
	return c.Id
}

func (c Company) GetName() string {
	return c.Name
}

func (c Company) GetCountry() string {
	return c.Country
}

func (c *Company) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("company name cannot be empty")
	}
	c.Name = name
	return nil
}

func (c *Company) SetCountry(country string) {
	c.Country = country
}

func main(){
	company := NewCompany("345", "Kikuu", "China")

	fmt.Printf("Company ID: %s\n", company.GetID())
	fmt.Printf("Company Name: %s\n", company.GetName())
	fmt.Printf("Company Country: %s\n", company.GetCountry())

	company.SetName("Google")
	fmt.Printf("Company Name: %s\n", company.GetName())



}