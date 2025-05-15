package assortedtoogo03smallestpositive

func smallestPositive(nums []int) int {
	var least int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			least = nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 && nums[i] < least {
			least = nums[i]
		}
	}

	return least
}
