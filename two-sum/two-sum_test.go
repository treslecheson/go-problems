package twosum

import (
	"reflect"
	"testing"
)

var nums1 = []int{2, 7, 11, 15}
var target1 = 9
var result1 = []int{0, 1}

var nums2 = []int{3, 2, 4}
var target2 = 6
var result2 = []int{1, 2}

var nums3 = []int{-1, -2, -3, -4, -5}
var target3 = -8
var result3 = []int{2, 4}

var nums4 = []int{10, -3, 7, 1, -6}
var target4 = 4
var result4 = []int{0, 4}

var nums5 = []int{3, 3}
var target5 = 6
var result5 = []int{0, 1}

var nums6 = []int{15, 2, 11, 7}
var target6 = 9
var result6 = []int{1, 3}

func TestTwoSum(t *testing.T) {
	assertEqual := func(t testing.TB, nums []int, target int, result []int) {
		got := twoSum(nums, target)
		wanted := result
		if reflect.DeepEqual(got, wanted) == false {
			t.Errorf("twoSum(%d, %d) = %d, want %d", nums, target, got, wanted)
		}
	}
	t.Run("test1", func(t *testing.T) {
		assertEqual(t, nums1, target1, result1)
	})
	t.Run("test2", func(t *testing.T) {
		assertEqual(t, nums2, target2, result2)
	})
	t.Run("negative test", func(t *testing.T) {
		assertEqual(t, nums3, target3, result3)
	})
	t.Run("positive and negative test", func(t *testing.T) {
		assertEqual(t, nums4, target4, result4)
	})
	t.Run("unsored test", func(t *testing.T) {
		assertEqual(t, nums5, target5, result5)
	})
}
