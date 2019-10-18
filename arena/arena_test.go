package arena

import "testing"

func TestArena(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Arena()
		})
	}
}

func TestFight(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fight(); got != tt.want {
				t.Errorf("Fight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOattack(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oattack(); got != tt.want {
				t.Errorf("Oattack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOhealth(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ohealth(); got != tt.want {
				t.Errorf("Ohealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oname(); got != tt.want {
				t.Errorf("Oname() = %v, want %v", got, tt.want)
			}
		})
	}
}

//test = ok
