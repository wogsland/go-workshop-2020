package main

import (
  "fmt"
  "errors"
)

type Thing string

type Person struct {
  name string
  age int
  favorite string
}

func main() {
  person := Person{
    name: "Ricco",
    age: 40,
    favorite: "blue",
  }
  name := Thing("Bradley")
  // var foo string
  fmt.Printf("nachoes are delicious! - %s\n", name)
  fmt.Printf("The teacher is %s\n", person.name)

  slice := make([]Person, 0, 10)
  fmt.Printf("length: %d\n", len(slice))

  slice = append(slice, Person{
    name: "Juan",
    age: 24,
    favorite: "rojo",
  })

  for i, p2 := range slice {
    fmt.Println(slice[i].format())
    p2.format()
  }

  bomb, err := fail()
  if err != nil {
    panic(err)
  }
  fmt.Println(bomb)
}

func (p Person) format() string {
  return "California"
}

func fail() (string, error) {
  return "bomb", errors.New("FAIL!")
}
