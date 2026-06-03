package raindrops

import "fmt"

func Convert(number int) string {
	var output string
	if !(number%3 == 0 || number%5 == 0 || number%7 == 0) {
		output = fmt.Sprintf("%d", number)
	}
	if number%3 == 0 {
		output = output + "Pling"
	}
	if number%5 == 0 {
		output = output + "Plang"
	}
	if number%7 == 0 {
		output = output + "Plong"
	}
	return output
}
