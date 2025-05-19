package assortedts01converttemperature

import "testing"

const farenheit1 float32 = 32
const celcius1 float32 = 0

const farenheit2 float32 = 212.0
const celcius2 float32 = 100.0

const farenheit3 float32 = 93.2
const celcius3 float32 = 34

func TestTofarenheit(t *testing.T) {
	assertEqual := func(t testing.TB, farenheit, celcius float32) {
		got := toFarenheit(celcius)
		wanted := farenheit

		if got != wanted {
			t.Errorf("Expected: %f, got: %f", wanted, got)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		assertEqual(t, farenheit1, celcius1)
	})
	t.Run("test 2", func(t *testing.T) {
		assertEqual(t, farenheit2, celcius2)
	})

}
