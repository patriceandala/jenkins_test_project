package printer

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hello(false)
		})
	}
}
