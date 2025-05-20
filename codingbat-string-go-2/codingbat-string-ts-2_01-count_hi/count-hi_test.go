package codingbatstringts201counthi

import "testing"

const text1 string = "hi"
const expected1 int = 1

const text2 string = "abc hi ho"
const expected2 int = 1

const text3 string = "HihiHI"
const expected3 int = 3

const text4 string = "hihi"
const expected4 int = 2

const text5 string = ""
const expected5 int = 0

const text6 string = "h"
const expected6 int = 0

func TestCountHi(t *testing.T) {
	assertEqual := func(t testing.TB, text string, expected int) {
		got := countHi(text)
		wanted := expected

		if got != wanted {
			t.Errorf("Expected: %d, got: %d", wanted, got)
		}
	}

	t.Run("test simple string", func(t *testing.T) {
		assertEqual(t, text1, expected1)
	})
	t.Run("tests small string", func(t *testing.T) {
		assertEqual(t, text2, expected2)
	})
	t.Run("tests string with capital HI", func(t *testing.T) {
		assertEqual(t, text3, expected3)
	})
	t.Run("tests string with no space", func(t *testing.T) {
		assertEqual(t, text4, expected4)
	})
	t.Run("tests empty string", func(t *testing.T) {
		assertEqual(t, text5, expected5)
	})
	t.Run("tests single character", func(t *testing.T) {
		assertEqual(t, text6, expected6)
	})

}
