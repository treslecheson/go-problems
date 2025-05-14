package assortedtoots01countduplicates

func countDuplicates(numbers []int, target int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			total++
		}

	}
	return total
}
