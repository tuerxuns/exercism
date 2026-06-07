package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleCount := 0
	sauceCount := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleCount += 50
		} else if layer == "sauce" {
			sauceCount += 0.2
		}
	}
	return noodleCount, sauceCount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	lastIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = lastIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(requiredAmount []float64, requiredPortions int) []float64 {
	scaledQuantities := []float64{}
	for _, amount := range requiredAmount {
		scaledQuantities = append(scaledQuantities, amount*float64(requiredPortions)/2)
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
