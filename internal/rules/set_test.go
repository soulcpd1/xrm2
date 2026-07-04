package rules

import "testing"

func TestDefaultSet(t *testing.T) {

	e := New()

	Default().Apply(e)

	if len(e.rules) == 0 {
		t.Fatal("rules not loaded")
	}
}
