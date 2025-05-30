package main

import "testing"

func TestIsAnagram(t *testing.T) {
	assertEqual := func(t testing.TB, str1, str2 string, wanted bool) {

		got := isAnagram(str1, str2)
		want := wanted

		if got != want {
			t.Errorf("got %v but wanted %v", got, want)
		}

	}
	t.Run("true case 1", func(t *testing.T) {
		assertEqual(t, "Sled", "LEDs", true)
	})
	t.Run("true case 2", func(t *testing.T) {
		assertEqual(t, "read", "dear", true)
	})
	t.Run("false case 1", func(t *testing.T) {
		assertEqual(t, "deer", "deere", false)
	})
	t.Run("true case 3", func(t *testing.T) {
		assertEqual(t, "listen", "Silent", true)
	})
	t.Run("false case 2", func(t *testing.T) {
		assertEqual(t, "hello", "bello", false)
	})
	t.Run("true case 4", func(t *testing.T) {
		assertEqual(t, "Triangle", "Integral", true)
	})
	t.Run("false different lengths", func(t *testing.T) {
		assertEqual(t, "abc", "ab", false)
	})
	t.Run("true case 2", func(t *testing.T) {
		assertEqual(t, "read", "dear", true)
	})
	t.Run("case sensitivity", func(t *testing.T) {
		assertEqual(t, "aA", "Aa", true)
	})
	t.Run("anagram with spaces", func(t *testing.T) {
		assertEqual(t, "rail safety", "fairy tales", true)
	})
	t.Run("anagram multi word", func(t *testing.T) {
		assertEqual(t, "the eyes", "they see", true)
	})
	t.Run("empty strings", func(t *testing.T) {
		assertEqual(t, "", "", true)
	})

}
