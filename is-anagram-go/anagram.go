package main

import "strings"

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if strings.Contains(strings.ToLower(str2), strings.ToLower(string(str1[i]))) {
			strings.Replace(str2, string(str1[i]), "", 1)
		} else {
			return false
		}
	}
	return true

}

func main() {
	println(isAnagram("Sled", "LEDs"))
	println(isAnagram("read", "dear"))
	println(isAnagram("deer", "deere"))
	println(isAnagram("read", "dear"))
}
