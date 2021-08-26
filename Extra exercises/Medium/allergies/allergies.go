// Package allergies determines whether they are allergic to a given item.
package allergies

var (
	allergyList = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
	allergyMap  = map[string]uint{
		"eggs":         1,
		"peanuts":      2,
		"shellfish":    4,
		"strawberries": 8,
		"tomatoes":     16,
		"chocolate":    32,
		"pollen":       64,
		"cats":         128,
	}
)

// Allergies returns a list of all allergies.
func Allergies(score uint) (output []string) {
	for i := 0; i < 8; i++ {
		if 1<<i&score != 0 {
			output = append(output, allergyList[i])
		}
	}

	return output
}

// AllergicTo gives a score and a substance, and returns true if allergic.
func AllergicTo(score uint, allergen string) (output bool) {
	return score&allergyMap[allergen] != 0
}
