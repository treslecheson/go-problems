package truncatesentencego

import "testing"

const s1 string = "Hello how are you Student"
const k2 int = 4

func TestTruncateString(t *testing.T) {
	assertEqual := func(t testing.TB, s string, k int, expected string) {
		got := truncateString(s, k)
		wanted := expected

		if got != wanted {
			t.Errorf("Expected: %s, got: %s", got, wanted)
		}

	}
	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, s1, k2)
	})
}
