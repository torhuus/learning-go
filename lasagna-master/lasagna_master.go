package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}

	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for i := 0; i < len(ingredients); i++ {
		switch ingredients[i] {
		case "noodles":
			noodles = noodles + 1
		case "sauce":
			sauce = sauce + 1
		}
	}
	return noodles * 50, sauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsRecipe, myRecipe []string) []string {
	lastIngredient := friendsRecipe[len(friendsRecipe)-1]

	myRecipe[len(myRecipe)-1] = lastIngredient

	return myRecipe
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(defaultPortion []float64, portions int) []float64 {
	var newIngredientAmounts []float64

	for i := 0; i < len(defaultPortion); i++ {
		newValue := (defaultPortion[i] / 2) * float64(portions)
		newIngredientAmounts = append(newIngredientAmounts, newValue)
	}

	return newIngredientAmounts
}
