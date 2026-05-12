package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutesPerLayer int) int {
	if minutesPerLayer == 0 {
		minutesPerLayer = 2
	}
	return len(layers) * minutesPerLayer
}

const noodlesPerLayer = 50
const saucePerLayer = 0.2

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += noodlesPerLayer
		case "sauce":
			sauce += saucePerLayer
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(theirs, ours []string) {
	ours[len(ours)-1] = theirs[len(theirs)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, portions int) []float64 {
	newRecipe := make([]float64, len(recipe))
	copy(newRecipe, recipe)

	scaleFactor := float64(portions) / 2
	for i, value := range newRecipe {
		newRecipe[i] = value * scaleFactor
	}
	return newRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
