package rules

import "testing"

func TestLoad(t *testing.T) {

	set, err := Load("/etc/xrm/rules.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if set == nil {
		t.Fatal("nil ruleset")
	}

	if len(set.Rules) == 0 {
		t.Fatal("empty ruleset")
	}
}
