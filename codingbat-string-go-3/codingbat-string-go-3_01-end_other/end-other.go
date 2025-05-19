package codingbatstringgo301endother

import "strings"

func endOther(str1, str2 string) bool {
	var longer string
	var shorter string

	if len(str1) > len(str2) {
		longer = str1
		shorter = str2
	} else {
		longer = str2
		shorter = str1
	}

	length := len(longer) - len(shorter)

	if strings.ToLower(longer[length:]) == strings.ToLower(shorter) {
		return true
	}
	return false

}
