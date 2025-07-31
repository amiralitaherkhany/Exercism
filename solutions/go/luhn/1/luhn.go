package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	str := strings.Replace(id, " ", "", -1)
	if len(str) <= 1 {
		return false
	}
	sum := 0
	sw := false
	for i := len(str) - 1; i > -1; i-- {
		v, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return false
		}
		if sw {
			if v2 := v * 2; v2 > 9 {
				sum += v2 - 9
			} else {
				sum += v2
			}
			sw = false
			continue
		}
		sum += v
		sw = true
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
