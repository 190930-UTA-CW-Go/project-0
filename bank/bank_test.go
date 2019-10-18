package bank

import "testing"

func TestAccount(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Account(); got != tt.want {
				t.Errorf("Account() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bank()
		})
	}
}

func TestPbfunds(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pbfunds(); got != tt.want {
				t.Errorf("Pbfunds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFail(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fail()
		})
	}
}

//test = ok
