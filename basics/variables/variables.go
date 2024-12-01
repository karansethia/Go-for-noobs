package main

import "fmt"

const myConstant = 245 // constants and variables can also be declared here but cant use the shorthand declarations

func main(){
  var str string = "Karan"
  // or 
  var str2 = "Sethia" // Type inference

  var num int = 24

  // var num2 float32 = 34.2

// SHorthand syntax
  str3 := "Mohan"

  fmt.Println(str, str2, num, str3)

  fmt.Println(myConstant)

const name string = "Karan" // constants in Go

  //constant groups or multiple constants
  const (
    port = 5000
    host = "localhost"
  )
  fmt.Println(port, host)

  number := 245
}
