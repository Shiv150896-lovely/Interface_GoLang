package main

import (
	"fmt"
	"log"
)

type Calculate interface {
	Add() float32
	Sub() int64
	Mul() float64
	Divide() float64
}

type Number struct {
	Radius1 float32
	Radius2 float32
}

func (n Number) Add() float32 {
	return n.Radius1 + n.Radius2
}

func (n Number) Sub() float32 {
	return n.Radius1 - n.Radius2
}

func (n Number) Mul() float32 {
	return n.Radius1 * n.Radius2
}

func (n Number) Divide() float32 {
	return n.Radius1 / n.Radius2
}

func main() {
	log.Println("add for two input.....starting")

	for {
		var choiceNumber string

		var input1 float32
		var input2 float32
		fmt.Println("enter input1 :- ")
		fmt.Scanln(&input1)
		fmt.Println("Enter the press +,-,*,/ :- ")
		fmt.Scanln(&choiceNumber)
		fmt.Println("enter input2 :- ")
		fmt.Scanln(&input2)
		r := Number{Radius1: input1, Radius2: input2}

		switch choiceNumber {
		case "+":
			fmt.Println("add Data :- ", r.Add())
		case "-":
			fmt.Println("sub Message :- ", r.Sub())
		case "*":
			fmt.Println("multiple Message :- ", r.Mul())
		case "/":
			fmt.Println("DIVIDE Message :- ", r.Divide())
		default:
			fmt.Println("invalid your choice operator and number.....")
		}

		fmt.Println("done calculate.")

		fmt.Println("do you want to perform calculate yes/no")
		var choice string
		fmt.Scanln(&choice)
		if choice != "yes" {
			break
		}
	}

	log.Println("Calculator finished.")
	fmt.Println("completed")

}

//-------------------simple interface for math package --------------------

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Area interface {
// 	Add() float64
// 	Sub() float64
// }

// Number struct {
// 	Radius float64
// }

// func (c Circle) Add() float64 {
// 	return float64(math.Pi * c.Radius * c.Radius)
// }

// func (c Circle) Sub() float64 {
// 	return float64(c.Radius - c.Radius)
// }

// // func (c Circle)
// func main() {
// 	fmt.Println("Enter your number :-")
// 	var input float64
// 	fmt.Scanln(&input)

// 	r := Circle{Radius: input}
// 	fmt.Println("add message:-", r.Add())
// 	fmt.Printf("subsruction message : %f\n", r.Sub())

// 	fmt.Println("hello interface ....")

// }

//-----------------------simple interface ----------------

// package main

// import "fmt"

// type Area interface {
// 	Add() float32
// 	Sub() float32
// }

// type Circle struct {
// 	Radius int64
// }

// func (c Circle) Add() float32 {
// 	return float32(c.Radius + c.Radius)
// }

// func (c Circle) Sub() float64 {
// 	return float64(c.Radius - c.Radius)
// }

// // func (c Circle)
// func main() {
// 	fmt.Println("Enter your number :-")
// 	var input int64
// 	fmt.Scanln(&input)

// 	r := Circle{input}
// 	fmt.Println("add message:-", r.Add())
// 	fmt.Printf("subsruction message : %f\n", r.Sub())

// 	fmt.Println("hello interface ....")

// }
