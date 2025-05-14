package assortedtoots01countduplicates

import (
	"math/rand/v2"
	"testing"
)

func nums() []int {
	length := rand.IntN(10)
	nums := []int{1}
	for i := 0; i < length; i++ {
		nums = append(nums, rand.IntN(100))
	}
	return nums
}

func tar(nums []int) int {
	length := len(nums)
	ran := rand.IntN(length)
	return nums[ran]
}

func want(nums []int, target int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			total++
		}

	}
	return total
}

func TestCountDuplicates(t *testing.T) {
	assertEqual := func(t testing.TB, numbers []int, target, expected int) {
		got := countDuplicates(numbers, target)
		wanted := expected

		if got != wanted {
			t.Errorf("Expected: %d, got: %d", wanted, got)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		numbers := nums()
		target := tar(numbers)
		assertEqual(t, numbers, target, want(numbers, target))
	})
	t.Run("test 2", func(t *testing.T) {
		numbers := nums()
		target := tar(numbers)
		assertEqual(t, numbers, target, want(numbers, target))
	})

	t.Run("test 3", func(t *testing.T) {
		numbers := nums()
		target := tar(numbers)
		assertEqual(t, numbers, target, want(numbers, target))
	})

	t.Run("test 4", func(t *testing.T) {
		numbers := nums()
		target := tar(numbers)
		assertEqual(t, numbers, target, want(numbers, target))
	})

}
