package main

import "testing"

func TestConvertAddress(t *testing.T) {
	assertEqual := func(t testing.TB, ip, wanted string) {
		got := convertAddress(ip)
		want := wanted

		if got != wanted {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	}
	t.Run("Returns defanged IP for a valid IP address", func(t *testing.T) {
		assertEqual(t, "192.168.1.9", "192[.]168[.]1[.]9")
	})
	t.Run("Returns defanged IP for a valid IP address", func(t *testing.T) {
		assertEqual(t, "10.0.0.1", "10[.]0[.]0[.]1")
	})

	t.Run("Returns empty string when an octet is greater than 255", func(t *testing.T) {
		assertEqual(t, "256.100.50.25", "")
	})
	t.Run("Returns empty string when an octet is greater than 255", func(t *testing.T) {
		assertEqual(t, "100.300.50.25", "")
	})
	t.Run("Returns empty string when an octet is greater than 255", func(t *testing.T) {
		assertEqual(t, "100.50.25.999", "")
	})

	t.Run("Handles edge values 0 and 255 correctly", func(t *testing.T) {
		assertEqual(t, "0.0.0.0", "0[.]0[.]0[.]0")
	})
	t.Run("Handles edge values 0 and 255 correctly", func(t *testing.T) {
		assertEqual(t, "255.255.255.255", "255[.]255[.]255[.]255")
	})

	t.Run("Returns empty string when an octet is negative", func(t *testing.T) {
		assertEqual(t, "-1.100.50.25", "")
	})
	t.Run("Returns empty string when an octet is negative", func(t *testing.T) {
		assertEqual(t, "100.-20.50.25", "")
	})
	t.Run("Returns empty string when an octet is negative", func(t *testing.T) {
		assertEqual(t, "100.50.-5.25", "")
	})
	t.Run("Returns empty string when an octet is negative", func(t *testing.T) {
		assertEqual(t, "100.50.25.-100", "")
	})

}
