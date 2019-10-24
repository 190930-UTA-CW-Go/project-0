package bank

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/artytheparty/project-0/emp"
	"github.com/artytheparty/project-0/usr"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func TestDeposit(t *testing.T) {
	//opening conncetion to the database
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	var tempUSR usr.User = GetUsrInfo("akhv", dataB)
	fmt.Println(tempUSR)
	//if (tempUSR.username) == "akhv" {

	//}
}

func TestRecoverError(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RecoverError()
		})
	}
}

func TestPrintUserInfo(t *testing.T) {
	type args struct {
		db       *sql.DB
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintUserInfo(tt.args.db, tt.args.username)
		})
	}
}

func TestWithdraw(t *testing.T) {
	type args struct {
		a    usr.User
		dAmt float64
		db   *sql.DB
	}
	tests := []struct {
		name string
		args args
		want usr.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Withdraw(tt.args.a, tt.args.dAmt, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsrInfo(t *testing.T) {
	type args struct {
		username string
		db       *sql.DB
	}
	tests := []struct {
		name string
		args args
		want usr.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsrInfo(tt.args.username, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsrInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEmployeeInfo(t *testing.T) {
	type args struct {
		usrname string
		db      *sql.DB
	}
	tests := []struct {
		name string
		args args
		want emp.Employee
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEmployeeInfo(tt.args.usrname, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployeeInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUserDB(t *testing.T) {
	type args struct {
		db *sql.DB
		a  usr.User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUserDB(tt.args.db, tt.args.a)
		})
	}
}

func TestCreateNewUserEntry(t *testing.T) {
	type args struct {
		username string
		password string
		fname    string
		lname    string
		db       *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewUserEntry(tt.args.username, tt.args.password, tt.args.fname, tt.args.lname, tt.args.db)
		})
	}
}

func TestCreateNewAccountEntry(t *testing.T) {
	type args struct {
		usrid   string
		acctype string
		bal     float64
		db      *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewAccountEntry(tt.args.usrid, tt.args.acctype, tt.args.bal, tt.args.db)
		})
	}
}
