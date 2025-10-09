package main
import "fmt"

func main() {
  if 203 > 178 {
    fmt.Println("203 is greater than 178")
  }

  a:= 203
  b:= 178
  if b < a {
    fmt.Println("b is less than b")
  }

  num := 1 
  if (num>=1){
	num+=5
	fmt.Println("Here is the new number", num)
  }else{
	fmt.Println("number must be positive")
  }

  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}