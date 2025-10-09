package main

import "fmt"

func main() {
  slice1 := []int{} //or var slice1 []int

  fmt.Println(len(slice1))
  fmt.Println(cap(slice1)) //returns the capacity of the slice
  fmt.Println(slice1)

  slice2 := []string{"hey", "hello", "greetings", "salut"}

  fmt.Println(len(slice2))
  fmt.Println(cap(slice2)) //returns the capacity of the slice
  fmt.Println(slice2)

  arr1 := [5]int{43,23,12,64,14} //or var arr1 [5]int{43,23,12,64,14}
  slice3 := arr1[1:3]

  fmt.Printf("myslice = %v\n", slice3)
  fmt.Printf("length = %d\n", len(slice3))
  fmt.Printf("capacity = %d\n", cap(slice3))

  arr := [5]int{1, 2, 3, 4, 5}

  slice4 := arr[1:4]    // [2, 3, 4] - shares arr's memory
  slice5 := arr[:3]     // [1, 2, 3] - from start to index 3
  slice6 := arr[2:]     // [3, 4, 5] - from index 2 to end

  slice4[0] = 99
  fmt.Println(arr) // [1, 99, 3, 4, 5]
  fmt.Println(slice4)
  fmt.Println(slice5)
  fmt.Println(slice6)

  //using make function
  
  s := make([]int, 0, 10)  // length 0, capacity 10

  // Appending
  s = append(s, 1) // [1]
  s = append(s, 2, 3, 4) // [1, 2, 3, 4]
  s = append(s, []int{5, 6}...) // [1, 2, 3, 4, 5, 6] or newSlice := []int{5, 6} so s = append(s, newSlice...)

  // Slicing
  s2 := s[1:3] // [2, 3]
  s3 := s[:2] // [1, 2]
  s4 := s[3:] // [4, 5, 6]

  fmt.Println(s2)
  fmt.Println(s3)
  fmt.Println(s4)

  // Copying
  s1 := []int{1, 2, 3, 4} // len=4, cap=4
  newSlice := []int{5, 6}  // len=2

  s = append(s1, make([]int, len(newSlice))...) //creates []int{0, 0} so current len of s is 6
  

  //copy(destination, source)
  copy(s[len(s)-len(newSlice):], newSlice) //copy(s[4:], newSlice)
  fmt.Println(s)



}