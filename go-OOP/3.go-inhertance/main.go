package main

import "fmt"

//inheritance using composition and interface

type Entity struct{
	Id string
	CreatedAt string
	UpdatedAt string
}

func NewEntity(id string) Entity{
	return Entity{
		Id: id,
		CreatedAt: "2024-01-01",
		UpdatedAt: "2024-01-01",
	}
}

func (e *Entity) SetId(id string) error{
	if (e.Id ==""){
		fmt.Println("The Entity null should not to be empy")
	}
	e.Id = id
	return nil //this means it returns no error
}

func (e Entity) GetID() string {
	return e.Id
}

func (e Entity) GetCreatedAt() string {
	return e.CreatedAt
}

func (e *Entity) UpdateTimestamp() {
	e.UpdatedAt = "2024-12-19"
}

func (e Entity) DisplayInfo() string {
	return fmt.Sprintf("ID: %s, Created: %s, Updated: %s", 
		e.Id, e.CreatedAt, e.UpdatedAt)
}

type Company struct {
	Entity    // this provides inheritance
	Name      string
	Country   string
	Employees int
}

func NewCompany(id, name string, country string, employees int) Company{
	return Company{
		Entity: NewEntity(id),
		Name: name,
        Employees: employees,
	}
}

func (c *Company) SetName(name string) error {
	if(c.Name == ""){
		fmt.Println("Company Name cannot be empty")
	}
	c.Name = name
	return nil

}

func (c Company) GetName() string {
	return c.Name
}

func (c *Company) SetEmployees(count int) {
	c.Employees = count
	c.UpdateTimestamp() // Can call parent method
}

func (c Company) GetEmployeeCount() int {
	return c.Employees
}

func main(){
	company := NewCompany("C123", "Kikuu", "China", 500)

	fmt.Printf("Inherited ID: %s\n", company.GetID())
	fmt.Printf("Created At inherited: %s\n", company.GetCreatedAt())
	fmt.Printf("Employees: %d\n", company.GetEmployeeCount()) 
}