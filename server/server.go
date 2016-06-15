package main

import(
  "fmt"
  "encoding/json"
)

type Person struct{
    Name string `json:"name"`
    Title string `json:"title"`
    Age int `json:"age"`
}

func main(){
  person := Person{"Brian", "Developer", 25}
  json, err := json.Marshal(person)
  if err != nil{
    fmt.Println("Error")
  }
  fmt.Println(string(json))
}
