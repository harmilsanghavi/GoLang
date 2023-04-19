package numerics

import (
	"fmt"
	"practice/userinput"
)

func AsciI() {
	var a string
	fmt.Printf("Enter First Value :---\n")
	fmt.Scanf("%s", &a)

	for i := 0; i < len(a); i++ {
		fmt.Printf("asccii value of %c is %d\n", a[i], a[i])
	}
}
func Average() {

	var c int
	a, b := userinput.User()
	fmt.Printf("Enter Third Value :---\n")
	fmt.Scan(&c)

	var sum = a + b + c
	fmt.Printf("sum is :---- %v\n", sum/3)
}
func OddEven() {
	var c int
	fmt.Printf("Enter Third Value :---\n")
	fmt.Scan(&c)

	if c%2 == 0 {
		fmt.Printf("%v is Even Number\n", c)
	} else {
		fmt.Printf("%v is Odd Number\n", c)
	}
}
func PositivieNegative() {
	var c int
	fmt.Printf("Enter Third Value :---\n")
	fmt.Scan(&c)

	if c > 0 {
		fmt.Printf("%v is Positive Number\n", c)
	} else if c == 0 {
		fmt.Printf("%v is 0 Number\n", c)
	} else {
		fmt.Printf("%v is Negative Number\n", c)
	}
}
func Swap() {

	a, b := userinput.User()

	fmt.Printf("Before Swap a value is %v and b value is %v \n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After Swap a value is %v and b value is %v \n", a, b)
}
func Add() {
	a, b := userinput.User()
	fmt.Printf("sum is :---- %v\n", a+b)
}
