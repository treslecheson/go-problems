package codingbatstringts201counthi

import (
	"strings"
)

func countHi(text string) int {
	count := 0
	leng := len(text)
	for i := 0; i < leng-1; i++ {
		secondIndex := i + 2
		segment := text[i:secondIndex]
		if strings.ToLower(segment) == "hi" {
			count++
		}
	}
	return count
}
