package emp

import (
	"reflect"
	"testing"
)

func TestCreateNewEmployee(t *testing.T) {
	type args struct {
		id       string
		username string
		pass     string
		fName    string
		lName    string
	}
	tests := []struct {
		name string
		args args
		want Employee
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateNewEmployee(tt.args.id, tt.args.username, tt.args.pass, tt.args.fName, tt.args.lName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNewEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetEmployeeID(t *testing.T) {
	tests := []struct {
		name string
		a    *Employee
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEmployeeID(); got != tt.want {
				t.Errorf("Employee.GetEmployeeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetEmployeeUsername(t *testing.T) {
	tests := []struct {
		name string
		a    *Employee
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEmployeeUsername(); got != tt.want {
				t.Errorf("Employee.GetEmployeeUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetEmployeePass(t *testing.T) {
	tests := []struct {
		name string
		a    *Employee
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEmployeePass(); got != tt.want {
				t.Errorf("Employee.GetEmployeePass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetEmployeeFName(t *testing.T) {
	tests := []struct {
		name string
		a    *Employee
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEmployeeFName(); got != tt.want {
				t.Errorf("Employee.GetEmployeeFName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetEmployeeLName(t *testing.T) {
	tests := []struct {
		name string
		a    *Employee
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEmployeeLName(); got != tt.want {
				t.Errorf("Employee.GetEmployeeLName() = %v, want %v", got, tt.want)
			}
		})
	}
}
