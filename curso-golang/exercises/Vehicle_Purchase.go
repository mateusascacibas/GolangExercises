package main

import (
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	array := []string{option1, option2}
	sort.Strings(array)
	return array[0] + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var porc = 0.0
	if age < 3 {
		porc = 0.8
	} else if age >= 10 {
		porc = 0.5
	} else {
		porc = 0.7
	}
	return originalPrice * porc

}
