package main

import ("fmt")

func main() {

  var map1 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}

  fmt.Printf("a\t%v\n", map1)

  var a = make(map[string]string) 
  a["Id"] = "234"
  a["Name"] = "Herman"
  a["Age"] = "23"
                                 
  b := make(map[string]int)
  b["One"] = 1
  b["Two"] = 2
  b["Three"] = 3
  b["Four"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)

  //create an empty map
  var x = make(map[string]string) //false
  var y map[string]string //true

  fmt.Println(x == nil)
  fmt.Println(y == nil)

  //access map by its element
  fmt.Println(a["Name"])

  //update and add map elements
  a["Age"] = "25"
  b["Five"] = 5

  fmt.Println(a)
  fmt.Println(b)

  //delete element from map
  delete(a,"Id")
  delete(b, "Three")
  fmt.Println(a)
  fmt.Println(b)

  //check elements from map
  val1, exst1 := a["Name"]
  val2, exst2 := b["Six"]
  _, exst3 := a["Id"]
  fmt.Println(val1, exst1)
  fmt.Println(val2, exst2)
  fmt.Println(exst3)

  //maps are reference
  c := b
  fmt.Println(c)
  c["Six"] = 6
  fmt.Println(c)

  //iterate over Maps
   j := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var k []string             
  k = append(k, "one", "two", "three", "four")

  for k, v := range a {        // loop with no order
    fmt.Printf("%v : %v, ", k, v)
  }

  fmt.Println()

  for _, element := range k {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, j[element])
  }
}