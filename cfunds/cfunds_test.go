package cfunds

import "testing"

func TestCfunds(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cfunds(); got != tt.want {
				t.Errorf("Cfunds() = %v, want %v", got, tt.want)
			}
		})
	}
}

//test = ok
