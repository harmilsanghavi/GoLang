package main

import (
	"fmt"
	"practice/arrayfunc"
	"practice/mathfunction"
	"practice/numerics"
	"practice/stringfunction"
)

func main() {
	var option int

	fmt.Printf("enter 1 for add 2 Number\n")
	fmt.Printf("enter 2 for average 3 Number\n")
	fmt.Printf("enter 3 for swap 2 Number\n")
	fmt.Printf("enter 4 for Scanf demo\n")
	fmt.Printf("enter 5 for get Ascii Value\n")
	fmt.Printf("enter 6 for get Absoluate Value\n")
	fmt.Printf("enter 7 for get Absoluate Value for integer\n")
	fmt.Printf("enter 8 for get Power of Number\n")
	fmt.Printf("enter 9 for get Power of Number without math.Pow\n")
	fmt.Printf("enter 10 for get max number using math.Max\n")
	fmt.Printf("enter 11 for get max number without using math.Max\n")
	fmt.Printf("enter 12 for get max number using math.Min\n")
	fmt.Printf("enter 13 for get max number without using math.Min\n")
	fmt.Printf("enter 14 for get math.ceil\n")
	fmt.Printf("enter 15 for get math.floor\n")
	fmt.Printf("enter 16 for get math.Mod\n")
	fmt.Printf("enter 17 for check given number is even or odd\n")
	fmt.Printf("enter 18 for check given number is Positive or Negative\n")
	fmt.Printf("enter 19 for dynamic array\n")

	fmt.Scan(&option)

	switch option {
	case 1:
		numerics.Add()
	case 2:
		numerics.Average()
	case 3:
		numerics.Swap()
	case 4:
		stringfunction.ScanF()
	case 5:
		numerics.AsciI()
	case 6:
		mathfunction.Absoulate()
	case 7:
		mathfunction.AbsoulateInt()
	case 8:
		mathfunction.PowerNumber()
	case 9:
		mathfunction.PowerWithoutPow()
	case 10:
		mathfunction.MaxNumber()
	case 11:
		mathfunction.MaxNumberWithoutMax()
	case 12:
		mathfunction.MinNumber()
	case 13:
		mathfunction.MinNumberWithoutMin()
	case 14:
		mathfunction.CeiL()
	case 15:
		mathfunction.FlooR()
	case 16:
		mathfunction.MoD()
	case 17:
		numerics.OddEven()
	case 18:
		numerics.PositivieNegative()
	case 19:
		arrayfunc.ArrayDynamic()
	default:
		fmt.Printf("Not Match\n")
	}
}
