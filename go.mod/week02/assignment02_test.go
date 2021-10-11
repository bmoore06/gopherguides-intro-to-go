package main

import "fmt"

func main() {
//array
  
//empty act variable, I was very confused as to what this actually is.
  
//Test File Naming
  
//exp variable
  exp := [4]string{"Ben", "Sam", "Becca", "Ollie"}
  for i := 0; i < len(exp); i++ {
    fmt.Println(exp[i])
  }

//Slices
  exp := []string{"Ben", "Sam", "Becca", "Ollie"}
  for i := 0; i < len(exp); i++ {
    fmt.Println(exp[i])
    fmt.Println(len(exp))
  }
  
//Maps
  exp := map[string]{string}
  exp["Ben"]
  exp["Sam"]
  exp["Becca"]
  exp["Ollie"]
  
  for key, value := range exp {
	  fmt.Printf("%s plays %s\n", key, value)
    fmt.Println(len(exp))
  key := ["Ben", "Sam", "Becca", "Ollie"]

  value, ok := exp[key]
  if !ok {
	  return fmt.Errorf("Key not found: %q", key) // Key not found: "foo"
}
  } 
   
  
