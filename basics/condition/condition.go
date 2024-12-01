package main

import "fmt"

func main() {
  // classic if else
  age := 12
  if age > 16 {
    fmt.Println("Adult")
  } else {
    fmt.Println("child")
  }

  role := "admin"
   hasPermission := true
  if role == "admin" || hasPermission {
    println("permission given")
  }

  if name:= "Karan"; name == "Mohan" {
    fmt.Println("This is Mohan ", name)
  } else {
    fmt.Println("This is not MOhan this is ", name)
  }


}
