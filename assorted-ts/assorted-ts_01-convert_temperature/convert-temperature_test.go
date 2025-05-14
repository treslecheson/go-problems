package assortedts01converttemperature

import "testing"

const farenheight1 float32 = 32
const celcius1 float32 = 0

const farenheight2 float32 = 212.0
const celcius2 float32 = 100.0

func TestToFarenheight(t *testing.T) {
	assertEqual := func(t testing.TB, farenheight, celcius float32) {
		got := toFarenheight(celcius)
		wanted := farenheight

		if got != wanted {
			t.Errorf("Expected: %f, got: %f", wanted, got)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, farenheight1, celcius1)
	})
	t.Run("test 2", func(t *testing.T) {
		assertEqual(t, farenheight2, celcius2)
	})

}
