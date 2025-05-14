package assortedts03findmax

import (
	"math/rand/v2"
	"testing"
)

func greater(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

var num1 int = rand.IntN(100)
var num2 int = rand.IntN(100)
var expected int = greater(num1, num2)

func TestFindMax(t *testing.T) {
	assertEqual := func(t testing.TB, num1, num2, expected int) {

		got := findMax(num1, num2)
		wanted := expected

		if got != wanted {
			t.Errorf("Expected: %d, got: %d", wanted, got)
		}
	}
	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, num1, num2, expected)
	})
}
