package stats

import "testing"

func TestPhealth(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Phealth(); got != tt.want {
				t.Errorf("Phealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPattack(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pattack(); got != tt.want {
				t.Errorf("Pattack() = %v, want %v", got, tt.want)
			}
		})
	}
}

//test = ok
