package main

import "fmt"

// TODO: define um 'OvenTime' constante
const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}

func main() {
	fmt.Println(RemainingOvenTime(2))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))
}
