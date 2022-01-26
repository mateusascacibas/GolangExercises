package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 2.3
	xs := fmt.Sprint(x)
	fmt.Print("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)
	fmt.Printf("O valor de x é %.2f", x)
}
