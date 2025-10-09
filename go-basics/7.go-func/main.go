package main

import "fmt"

func MyMessage() {
	fmt.Println("Hello Word")
}

//parameters & argumnets
func MyName(fullName string, age int8){
	fmt.Println("My name is", fullName, "and age is", age)
}

//return values
func NumAddition(x int8, y int8) int8{
	return x+y
}

func NumDivision(x int8, y int8) int8{
	return x/y
}

//multiple return
func ConstPI(pi float32, a string) (result float32, txt string){
	result = pi
	txt = a + "PI value"
	return
}

//recursion
func TestCount(x int8) int8 {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return TestCount(x + 1)
}

func NumFactorial(x int64) (y int64){
	if (x>0){
		y= x * NumFactorial(x-1)
	}else{
		y=1
	}
	return
}


func main(){
    
	MyMessage()

	MyName("Herman", 23)

	fmt.Println("Addition of two numbers",NumAddition(5,9))

	div := NumDivision(81,9)
	fmt.Println("Division of two numbers",div)

	fmt.Println(ConstPI(3.14, "is "))

	a, b := ConstPI(3.24, "is not ")
	fmt.Println(a,b)

	_, x := ConstPI(1.24, "1.24 won't be displayed by the passing argument however 1.24 is not ")
	fmt.Println(x)

	TestCount(1)

	fmt.Println(NumFactorial(1))

}