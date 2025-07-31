package isogram

import "strings"

func IsIsogram(word string) bool {
	hash := make(map[string]string)
	for _, c := range word {
		str := strings.ToLower(string(c))
		if str == "-" || str == " " {
			continue
		}
		if _, ok := hash[str]; ok {
			return false
		}
		hash[str] = str
	}
	return true
}
