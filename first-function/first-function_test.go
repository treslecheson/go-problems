package firstfunction

import "testing"

const num1 int = 10
const expected1 int = 5

const num2 int = 34
const expected2 int = 17

func TestHalf(t *testing.T) {
	assertEqual := func(t testing.TB, num, expected int) {
		got := Half(num)
		wanted := expected

		if got != wanted {
			t.Errorf("Expected: %d, got: %d", got, wanted)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, num1, expected1)
	})
	t.Run("test 2", func(t *testing.T) {
		assertEqual(t, num2, expected2)
	})

}
