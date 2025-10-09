package main
import ("fmt")

type Student struct {
  name string
  age int8
  sex string
  course string
}

func main() {
  var st1 Student

  st1.name = "Hege"
  st1.age = 19
  st1.sex = "Male"
  st1.course = "Go-Lang Programming Language"

  printPerson(st1)
}

func printPerson(st Student) {
  fmt.Println("Name: ", st.name)
  fmt.Println("Age: ", st.age)
  fmt.Println("Sex: ", st.sex)
  fmt.Println("Course: ", st.course)
}