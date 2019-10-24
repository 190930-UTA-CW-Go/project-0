package usr

import (
	"reflect"
	"testing"

	"github.com/artytheparty/project-0/acc"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		id       string
		username string
		password string
		fName    string
		lName    string
		accounts []acc.Account
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.id, tt.args.username, tt.args.password, tt.args.fName, tt.args.lName, tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetAccounts(t *testing.T) {
	tests := []struct {
		name string
		a    *User
		want []acc.Account
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetAccounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_PrintAccounts(t *testing.T) {
	tests := []struct {
		name string
		a    *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.PrintAccounts()
		})
	}
}

func TestUpdateUserAccounts(t *testing.T) {
	type args struct {
		u User
		a []acc.Account
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateUserAccounts(tt.args.u, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUserAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsrID(t *testing.T) {
	tests := []struct {
		name string
		a    *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetUsrID(); got != tt.want {
				t.Errorf("User.GetUsrID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsrUsername(t *testing.T) {
	tests := []struct {
		name string
		a    *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetUsrUsername(); got != tt.want {
				t.Errorf("User.GetUsrUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsrPassword(t *testing.T) {
	tests := []struct {
		name string
		a    *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetUsrPassword(); got != tt.want {
				t.Errorf("User.GetUsrPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsrFName(t *testing.T) {
	tests := []struct {
		name string
		a    *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetUsrFName(); got != tt.want {
				t.Errorf("User.GetUsrFName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsrLName(t *testing.T) {
	tests := []struct {
		name string
		a    *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetUsrLName(); got != tt.want {
				t.Errorf("User.GetUsrLName() = %v, want %v", got, tt.want)
			}
		})
	}
}
