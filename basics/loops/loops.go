package main

import "fmt"

func main(){
  //while loop as a for loop construct
  i := 0
  for i <= 4 {
    fmt.Println(i)
    i++
  }

  // infinite loop
  // for {
  //   fmt.Println("hello")
  // }

  // classic for loop
  for i := 0; i <= 4; i++ {
    if i == 2 {
      continue
    }
    if i == 3 {
      break
    }
    fmt.Println(i)
  }

  // loop in range
  for i := range 3 {
    fmt.Println(i)
  }
  
}
