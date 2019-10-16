package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "os"
	"project-0/employee"
	_ "project-0/employee"
	"text/tabwriter"

	_ "github.com/lib/pq"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//db.Exec("INSERT INTO pokemon VALUES (4, 'Eevee')")
	//db.Exec("INSERT INTO pokemon VALUES (8, 'tyrannitar')")
	//	getAll(db)
	//GetAll3(db) //tickets

	//se(db, 2)
	//employee.Approvedeny()
	//searchByName(db, "Eeesevee")
	//employee.NewAcc()
	//GetAll4(db) //admins
	//employee.SearchUser("thirdacc")
	employee.Welcome()
	//GetAll2(db)
	//GetAll22(db)
	//GetAll3(db)
	//GetAll4(db)
	//getAll(db)
	//employee.ManagerLogin()
	////db.Exec("INSERT INTO employeeAccounts VALUES ('adf', 'Eeeevee')")
	//GetAll3(db)
	//fmt.Println("die")
	//SearchByName2(db, "password")
	//employee.Welcome()
	//GetAll2(db) //users
	//GetAll22(db)
	//employee.Welcome()

}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM POKEMON")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM pokemon WHERE name = $1", searchvalue)
	var id int
	var name string
	row.Scan(&id, &name)
	fmt.Println(id, name)
}

func se(db *sql.DB, searchvalue int) {
	row := db.QueryRow("SELECT * FROM tickets WHERE ticketNum = $1", searchvalue)
	var u1 int
	var u2 string
	var u3 string
	var u4 string
	var u5 float32
	var u6 string
	var u7 string
	row.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
	fmt.Println(u1, u2, u3, u4, u5, u6, u7)
}

//GetAll2 comment
func GetAll2(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM EMPLOYEELOGIN")
	for rows.Next() {
		var u1 int
		var u2 string
		var u3 string
		var u4 string
		var u5 string
		rows.Scan(&u1, &u2, &u3, &u4, &u5)
		fmt.Println(u1, u2, u3, u4, u5)
	}
}

//GetAll3 suck
func GetAll3(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM TICKETS")
	for rows.Next() {
		var u1 int
		var u2 string
		var u3 string
		var u4 string
		var u5 float32
		var u6 string
		var u7 string
		rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
		fmt.Println(u1, u2, u3, u4, u5, u6, u7)
	}
}

//GetAll4 comment
func GetAll4(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM EMPLOYEEACCOUNTS")
	for rows.Next() {
		var u2 string
		var u3 string
		rows.Scan(&u2, &u3)
		fmt.Println(u2, u3)
	}
}

//GetAll22 comment
func GetAll22(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM EMPLOYEELOGIN")
	w := new(tabwriter.Writer)
	var fullstring string
	w.Init(os.Stdout, 0, 8, 6, '\t', 0)
	fmt.Fprintln(w, "USERNAME\tPASSWORD\tFIRSTNAME\tLASTNAME.")
	for rows.Next() {
		var u1 int
		var u2 string
		var u3 string
		var u4 string
		var u5 string
		rows.Scan(&u1, &u2, &u3, &u4, &u5)
		fullstring = (u2 + "\t" + u3 + "\t" + u4 + "\t" + u5 + "\t.")
		fmt.Fprintln(w, fullstring)
	}
	fmt.Fprintln(w)
	w.Flush()
}
