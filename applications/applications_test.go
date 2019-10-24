package applications

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestCreateAccountHolder(t *testing.T) {
	hold := CreateAccountHolder("username", "password", "fname", "lname", "s", 5000)
	if hold.username == "username" && hold.password == "password" && hold.fname == "fname" && hold.lname == "lname" && hold.types == "s" && hold.bal == 5000 {
		println("Success")
	} else {
		println("Error Occured")
	}
}
func TestAccountHolderInfo2(t *testing.T) {
	hold := CreateAccountHolder("username", "password", "fname", "lname", "s", 5000)
	if hold.AccountHolderInfo2() == "USERNAME: username PASSWORD: password FNAME: fname LNAME: lname TYPE: s BALANCE: 5000.00" {
		println("Success")
	} else {
		println("Error Occured")
	}
}

func TestAccountHolder_AccountHolderInfo2(t *testing.T) {
	type fields struct {
		username string
		password string
		fname    string
		lname    string
		types    string
		bal      float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AccountHolder{
				username: tt.fields.username,
				password: tt.fields.password,
				fname:    tt.fields.fname,
				lname:    tt.fields.lname,
				types:    tt.fields.types,
				bal:      tt.fields.bal,
			}
			if got := a.AccountHolderInfo2(); got != tt.want {
				t.Errorf("AccountHolder.AccountHolderInfo2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteToFile(t *testing.T) {
	type args struct {
		c []AccountHolder
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteToFile(tt.args.c)
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name string
		want []AccountHolder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintHolder(t *testing.T) {
	type args struct {
		a []AccountHolder
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintHolder(tt.args.a)
		})
	}
}

func TestAskNewAccount(t *testing.T) {
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
			AskNewAccount(tt.args.db)
		})
	}
}

func TestApproveAndAddToDB(t *testing.T) {
	type args struct {
		db *sql.DB
		a  []AccountHolder
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApproveAndAddToDB(tt.args.db, tt.args.a)
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		s []AccountHolder
		i int
	}
	tests := []struct {
		name string
		args args
		want []AccountHolder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
