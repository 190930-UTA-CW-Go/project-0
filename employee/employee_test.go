package employee

import (
	"database/sql"
	"fmt"
	_ "fmt" //wont
	"log"
	_ "log" //u
	"testing"

	_ "testing" //why
)

//OPENTEST opens database; needed to clean up code
func OPENTEST() *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	return db
}

// TestNewAcc Opens prompt to create a new account.
func TestNewAcc(t *testing.T) {

	userName := "thirdacc"
	password := "pass"
	fname := "DJ"
	lname := "Khaled"
	//db := OPENTEST()
	//defer db.Close()
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	/*db.Exec("INSERT INTO employeeLogin(userName,password,fname, lname)"+
	"VALUES($1,$2,$3, $4)", userName, password, fname, lname)*/
	db.Exec("INSERT INTO employeeLogin(userName,password,fname, lname)"+
		"VALUES($1,$2,$3, $4)", userName, password, fname, lname)
	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", userName)
	var u1 int
	var u2 string
	var u3 string
	var u4 string
	var u5 string
	row.Scan(&u1, &u2, &u3, &u4, &u5)
	fmt.Println(userName, password, fname, lname)
	fmt.Println(u2, u3, u4, u5)
	if userName == u2 && password == u3 && fname == u4 && lname == u5 {
		log.Printf("TestNewAcc passed")
	} else {
		log.Fatal("TestNewAcc failed. Inserted userName/password/fname/lname values: ",
			userName, password, fname, lname,
			". Retrieved values are: ", u2, u3, u4, u5)
	}
}
