package userinput

import "fmt"

func User() (int, int) {
	var a, b int
	fmt.Printf("Enter First Value :---\n")
	fmt.Scanf("%d", &a)

	fmt.Printf("Enter Second Value :---\n")
	fmt.Scanf("%d", &b)

	return a, b
}
