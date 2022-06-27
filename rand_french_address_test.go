package fake

import (
	"testing"
)

func TestFrenchAddressGenerator(t *testing.T) {
	t.Parallel()

	t.Run("should work with configuration", func(t *testing.T) {
		addr := PostalAddressString(&FrenchAddressGenerator{
			FrenchCities:  []string{"Codeville"},
			FrenchStreets: []string{"Rue des Tests Unitaires"},
			PostalCodes:   []string{"10000"},
			StreetNumbers: []string{"1"},
		})
		want := "1 Rue des Tests Unitaires, 10000 Codeville, France"
		if addr != want {
			t.Fatalf("unexpected address string, want: %#v, got: %#v", want, addr)
		}
	})

	t.Run("should work without configuration", func(t *testing.T) {
		// todo: improve
		addr := PostalAddressString(&FrenchAddressGenerator{})
		if len(addr) <= 9 {
			t.Fatalf("address is too short to be valid: %#v", addr)
		}
	})
}
