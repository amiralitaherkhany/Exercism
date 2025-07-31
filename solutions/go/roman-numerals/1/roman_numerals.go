package romannumerals

import "errors"
func ToRomanNumeral(input int) (string, error) {
    if input <= 0 || input >=4000{
		return "", errors.New("")
	}
	romanNumber := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	romanChars := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}
	result := ""

	for i, v := range romanNumber {
		for input >= v {
			result += romanChars[i]
			input = input - v
		}
	}

	return result, nil
}
