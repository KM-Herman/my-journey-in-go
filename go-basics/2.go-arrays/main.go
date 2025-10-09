package main

import "fmt"

var arr1 = [5]int{54, 34, 3, 55, 23} // outside the function

func main() {
	arr2 := [5]string{"wewe", "rere", "tete", "cece", "zeze"} //inside the function
	arr3 := [5]int{56,65} //partially initialized
    arr4 := [...]int{13,26,33,64,95} //fully initialized
	arr5 := [5]int{} //not initialized
	arr6 := [5]int{1:10,2:40} //initialize only specific elements

	fmt.Println(arr2)
	fmt.Println(arr1)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
}
