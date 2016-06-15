package main

import(
  "fmt"
  "encoding/json"
)

// Person struct
type Person struct{
    Name string `json:"name"`
    Title string `json:"title"`
    Age int `json:"age"`
}

func main(){
  // Assign a new instance of Person with grouped fields
  person := Person{"Brian", "Developer", 25}
  // Encode the person variable to json bytes
  json, err := json.Marshal(person)
  if err != nil{
    fmt.Println("Error")
  }
  // Convert json byte data into string and print
  fmt.Println(string(json))
}
