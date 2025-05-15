package assortedtoogo03smallestpositive

import "testing"

func TestSmallestPositive(t *testing.T) {
	assertEqual := func(t testing.TB, input []int, expect int) {
		got := smallestPositive(input)
		want := expect

		if got != want {
			t.Errorf("Got %d but was expecting %d", got, want)
		}
	}
	t.Run("test1", func(t *testing.T) {
		input1 := []int{10, 8, 22, 54, 9, 1}
		expect1 := 1
		assertEqual(t, input1, expect1)
	})
	t.Run("test2", func(t *testing.T) {
		input2 := []int{13, 43, -2, -19, 64, 3}
		expect2 := 3
		assertEqual(t, input2, expect2)
	})
	t.Run("test3", func(t *testing.T) {
		input3 := []int{-3, -1, -64, -10, 13, -12, 42, -1}
		expect3 := 13
		assertEqual(t, input3, expect3)
	})

}
