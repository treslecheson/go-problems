package main

import "testing"

const test1String1 string = "aB"
const test1String2 string = "BA"
const test1want bool = true

func TestIsAnagram(t *testing.T) {
	assertEqual := func(t testing.TB, str1, str2 string, wanted bool) {

		got := isAnagram(str1, str2)
		want := wanted

		if got != want {
			t.Errorf("got %v but wanted %v", got, want)
		}

	}
	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, test1String1, test1String2, test1want)
	})

}
