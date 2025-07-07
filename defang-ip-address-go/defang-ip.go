package main

import (
	"strconv"
	"strings"
)

func convertAddress(ip string) string {
	var defang string
	startSlice := 0

	if strings.Contains(ip, "-") {
		return ""
	}

	for i := 0; i < len(ip); i++ {
		// endSlice := strings.Index(ip, ".")

		octlet, err := strconv.Atoi(ip[startSlice:strings.Index(ip, ".")])
		if err != nil {
			return ""
		} else if 0 > octlet > 255 {
			return ""
		}
		println(octlet)

	}

	return defang
}

func main() {
	convertAddress("100.300.50.25")
}
