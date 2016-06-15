package main

import(
  "fmt"
  "github.com/gorilla/mux"
  "encoding/json"
  "net/http"
  "github.com/codegangsta/negroni"
)

// Person struct
type Person struct{
    Name string `json:"name"`
    Title string `json:"title"`
    Age int `json:"age"`
}

func main(){
  mux := mux.NewRouter()
  mux.HandleFunc("/", IndexHandler).Methods("GET")
  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":9090") // http://localhost:9090
}


func IndexHandler(w http.ResponseWriter, r *http.Request){
  // Assign a new instance of Person with grouped fields
  person := Person{"Brian", "Developer", 25}
  // Encode the person variable to json bytes
  json, err := json.Marshal(person)
  if err != nil{
    fmt.Println("Error")
  }
  // Convert json byte data into string and print
  w.Header().Set("Content-Type", "application/json")
  w.Write(json)
}
