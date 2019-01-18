package allergies

var allergyToScore = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(score uint) []string {
	var result []string
	// Using a list of strings to ensure that the order of the checks is consistent
	for _, substance := range []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"} {
		if AllergicTo(score, substance) {
			result = append(result, substance)
		}
	}
	return result
}

func AllergicTo(score uint, substance string) bool {
	substanceScore := allergyToScore[substance]
	return (score & substanceScore) == substanceScore
}
