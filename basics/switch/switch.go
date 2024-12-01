package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Karan"
	// classic switch
	switch name {
	case "Karan":
		fmt.Println("Karan is a great person")
	case "Mohan":
		fmt.Println("Mohan is a great person")
	default:
		fmt.Println("SOmeone else")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Print("oh not its workday")
	}

  // type switch
  getVariableType := func(i interface{}){ // interface{} is like any type of typescript
    switch i.(type) { // i.(type) is a special syntax that can only work in type switch
    case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("i is a string")
	case bool:
		fmt.Println("i is a bool")
	default:
		fmt.Println("i is of an unknown type")
    }
  }

  getVariableType(23)
  getVariableType("Hello")

}
