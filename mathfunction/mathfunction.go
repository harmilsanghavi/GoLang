package mathfunction

import (
	"fmt"
	"math"
	"practice/userinput"
)

func Absoulate() {
	var a float64
	fmt.Printf("Enter First Value :---\n")
	fmt.Scanf("%f", &a)

	fmt.Printf("absoulate value of %v Number is %v\n", a, math.Abs(a))
}

func AbsoulateInt() {
	var a int
	fmt.Printf("Enter First Value :---\n")
	fmt.Scanf("%d", &a)

	fmt.Printf("absoulate value of %v Number is %v\n", a, math.Abs(float64(a)))
}

func MinNumber() {
	a, b := userinput.User()

	fmt.Printf("smallest among %v %v is %v\n", a, b, math.Min(float64(a), float64(b)))
}
func MinNumberWithoutMin() {
	a, b := userinput.User()

	if a < b {
		fmt.Printf("smallest among %v %v is %v\n", a, b, a)
	} else {
		fmt.Printf("smallest among %v %v is %v\n", a, b, b)
	}
}

func MaxNumber() {
	a, b := userinput.User()

	fmt.Printf("largest among %v %v is %v\n", a, b, math.Max(float64(a), float64(b)))
}
func MaxNumberWithoutMax() {
	a, b := userinput.User()

	if a > b {
		fmt.Printf("largest among %v %v is %v\n", a, b, a)
	} else {
		fmt.Printf("largest among %v %v is %v\n", a, b, b)
	}
}

func PowerNumber() {
	a, b := userinput.User()

	fmt.Printf("power of %v is %v\n", a, math.Pow(float64(a), float64(b)))
}
func PowerWithoutPow() {
	a, b := userinput.User()
	var sum = 1
	for i := 1; i <= b; i++ {
		sum = sum * a
	}
	fmt.Printf("power of %v is %v\n", a, sum)
}
func CeiL() {
	var a float64
	fmt.Printf("Enter Third Value :---\n")
	fmt.Scan(&a)

	fmt.Printf("Ceil of %v is %v\n", a, math.Ceil(a))
}
func FlooR() {
	var a float64
	fmt.Printf("Enter Third Value :---\n")
	fmt.Scan(&a)

	fmt.Printf("Floor of %v is %v\n", a, math.Floor(a))
}
func MoD() {
	a, b := userinput.User()

	fmt.Printf("mod of %v %v is %v\n", a, b, math.Mod(float64(a), float64(b)))
}
