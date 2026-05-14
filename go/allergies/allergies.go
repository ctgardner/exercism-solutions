package allergies

var bitmasks = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	allergens := []string{}
	for allergen := range bitmasks {
		if AllergicTo(allergies, allergen) {
			allergens = append(allergens, allergen)
		}
	}
	return allergens
}

func AllergicTo(allergies uint, allergen string) bool {
	allergies = allergies & 255 // discard allergy values > 128
	mask := bitmasks[allergen]
	return allergies&mask == mask
}
