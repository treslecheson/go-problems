package assortedtoogo02removeduplicates

func contains(arr []string, target string) bool {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func removeDuplicates(arr []string) []string {
	new_arr := []string{arr[0]}

	for i := 1; i < len(arr); i++ {
		if contains(new_arr, arr[i]) == true {
			continue
		} else {
			new_arr = append(new_arr, arr[i])
		}
	}

	return new_arr
}
