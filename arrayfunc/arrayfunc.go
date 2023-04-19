package arrayfunc

import (
	"fmt"
)

func ArrayDynamic() {
	var n int

	fmt.Printf("enter no of element\n")
	fmt.Scan(&n)

	arr := []int{}

	var s int
	d := n
	for i := 0; i < n; i++ {
		fmt.Printf("enter %v value\n", d)
		fmt.Scan(&s)
		arr = append(arr, s)
		d--
	}
	fmt.Printf("Array value\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%v\n", arr[i])
	}
	var val string
	fmt.Printf("if you want to sort array in ascending press a and in decending press d and no to dont want\n")
	fmt.Scan(&val)

	if val == "a" {
		var temp int
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				if arr[j] < arr[i] {
					temp = arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
		fmt.Printf("Array after sort\n")
		for i := 0; i < n; i++ {
			fmt.Printf("%v\n", arr[i])
		}
	} else if val == "d" {
		var temp int
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				if arr[j] > arr[i] {
					temp = arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
		fmt.Printf("Array after sort\n")
		for i := 0; i < n; i++ {
			fmt.Printf("%v\n", arr[i])
		}
	} else {
		fmt.Printf("thank you\n")
	}
}
