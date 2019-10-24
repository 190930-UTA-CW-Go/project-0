package ui

import (
	"database/sql"
	"testing"

	"github.com/artytheparty/project-0/usr"
)

func TestMenu(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Menu(tt.args.db)
		})
	}
}

func TestUserMenu(t *testing.T) {
	type args struct {
		a  usr.User
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserMenu(tt.args.a, tt.args.db)
		})
	}
}

func TestEmployeeMenu(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmployeeMenu(tt.args.db)
		})
	}
}
