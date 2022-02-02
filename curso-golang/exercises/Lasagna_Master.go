package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * time
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := 0.0
	noodles := 0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	size := len(friendsList) - 1
	for i := 0; i < len(friendsList); i++ {
		if i == size {
			myList[i+1] = friendsList[i]
		}
	}
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, portions int) []float64 {
	test := 0.0
	if portions == 2 {
		test = 2.0
	} else {
		test = float64(portions) / 2.0
	}

	var scaled []float64
	if len(slice) != 0 {
		scaled = make([]float64, len(slice))
		for i := 0; i < len(slice); i++ {
			(scaled)[i] = (slice)[i] * float64(test)
		}
	}
	return (scaled)
}
