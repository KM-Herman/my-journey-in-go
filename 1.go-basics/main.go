package main
import "fmt"

const PI = 3.14 // or const PI int = 3.14 as typed constant

var i int = 10 //explicity type
var a = "Hey" //inference type

var j string = "Hello"

var age int8 = 25          // 1 byte (-128 to 127) - small numbers
var population int64 = 8000000000 // 8 bytes - very large numbers
var price float32 = 19.99  // 4 bytes - basic decimals
var scientificValue float64 = 3.141592653589793 // 8 bytes - high precision

func main(){
	fmt.Println("Hello Word")

	b := "9.3" //only inside the function

	fmt.Println(i)
	fmt.Println(a)
	fmt.Println(b)

	//for memory effeciency
	fmt.Printf("the value of age: %v and it's type: %T", age,age)
	fmt.Printf("the value of population: %#v and it's type: %T", population,population)
	fmt.Printf("the value of price: %v and it's type: %T", price,price)
	fmt.Printf("the value of scientificValue %v and it's type: %T", scientificValue,scientificValue)


	//%v is used to print the value of the arguments
	//%T is used to print the type of the arguments
	//%#v is used to print the value in go syntax format
	fmt.Printf("j has value: %v and type: %T", j, j)

	//constant
    fmt.Println("PI: ", PI)

}