package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }

  //continue
  for i:=0; i < 8; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }

  //break
  for i:=0; i < 10; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }

  //nested
  num := [2]string{"one", "two"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(num); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(num[i],fruits[j])
    }
  }

  //range
  
  colors := [3]string{"red", "orange", "green"}

  for idx,_ := range colors {
     fmt.Printf("%v\n", idx) // 0, 1, 2
  }

  for _, val := range colors {
     fmt.Printf("%v\n", val) //red, orange, green
  }
}