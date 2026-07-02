package routing

import "testing"

func TestMatcher(t *testing.T) {

	table := DefaultTable()

	tests := []struct {
		host string
		want string
	}{
		{"ya.ru", "direct"},
		{"vk.ru", "direct"},
		{"192.168.1.1", "direct"},
		{"1.1.1.1", "nebula"},
		{"google.com", "nebula"},
	}

	for _, tt := range tests {

		got := table.Match(tt.host)

		if got.Outbound != tt.want {
			t.Fatalf("%s => %s != %s", tt.host, got.Outbound, tt.want)
		}
	}
}
