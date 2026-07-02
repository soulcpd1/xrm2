package outbound

import "testing"

func TestDrivers(t *testing.T) {

	m := Default()

	if _, err := m.Driver("direct"); err != nil {
		t.Fatal(err)
	}

	if _, err := m.Driver("nebula"); err != nil {
		t.Fatal(err)
	}
}
