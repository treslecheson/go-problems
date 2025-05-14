package assortedts02calculateaverage

import (
	"math/rand/v2"
	"testing"
)

var num1 int = rand.IntN(100)
var num2 int = rand.IntN(100)
var num3 int = rand.IntN(100)
var expected float32 = float32((num1 + num2 + num3) / 3)

func TestCalculateAvg(t *testing.T) {
	assertEqual := func(t testing.TB, num1, num2, num3 int, expected float32) {
		got := calculateAvg(num1, num2, num3)
		wanted := expected

		if got != wanted {
			t.Errorf("Expected: %f, got: %f", wanted, got)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, num1, num2, num3, expected)
	})

}
