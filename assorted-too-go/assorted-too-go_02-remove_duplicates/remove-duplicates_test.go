package assortedtoogo02removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	assertEqual := func(t testing.TB, input []string, expected []string) {
		got := removeDuplicates(input)
		wanted := expected

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got: %v  Wanted %v", got, wanted)
		}
	}
	t.Run("test 1", func(t *testing.T) {
		arr1 := []string{"1", "1", "1", "1"}
		expected1 := []string{"1"}
		assertEqual(t, arr1, expected1)
	})
	t.Run("test 2", func(t *testing.T) {
		arr2 := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}
		expected2 := []string{"h", "e", "l", "o", "w", "r", "d"}
		assertEqual(t, arr2, expected2)
	})
	t.Run("test 3", func(t *testing.T) {
		arr3 := []string{"n", "o", "d", "u", "p", "e", "s"}
		expected3 := []string{"n", "o", "d", "u", "p", "e", "s"}
		assertEqual(t, arr3, expected3)
	})
}
