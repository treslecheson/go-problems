package codingbatstringgo301endother

import "testing"

const test1String1 string = "abc"
const test1String2 string = "abc"
const expected1 bool = true

const test2String1 string = "hiabc"
const test2String2 string = "abc"
const expected2 bool = true

const test3String1 string = "Hiabc"
const test3String2 string = "abcd"
const expected3 bool = false

const test4String1 string = "AbC"
const test4String2 string = "HiaBc"
const expected4 bool = true

const test5String1 string = "hello"
const test5String2 string = "world"
const expected5 bool = false

const test6String1 string = "abcXYA"
const test6String2 string = "abcCDG"
const expected6 bool = false

func TestEndOther(t *testing.T) {
	assertEqual := func(t testing.TB, str1, str2 string, expected bool) {
		got := endOther(str1, str2)
		want := expected
		if got != expected {
			t.Errorf("Got %t but was expecting %t", got, want)
		}
	}
	t.Run("same string", func(t *testing.T) {
		assertEqual(t, test1String1, test1String2, expected1)
	})
	t.Run("check simple string", func(t *testing.T) {
		assertEqual(t, test2String1, test2String2, expected2)
	})
	t.Run("check slightly different string", func(t *testing.T) {
		assertEqual(t, test3String1, test3String2, expected3)
	})
	t.Run("check case sensitivity", func(t *testing.T) {
		assertEqual(t, test4String1, test4String2, expected4)
	})
	t.Run("check different strings", func(t *testing.T) {
		assertEqual(t, test5String1, test5String2, expected5)
	})
	t.Run("check beginning of string", func(t *testing.T) {
		assertEqual(t, test6String1, test6String2, expected6)
	})
}
