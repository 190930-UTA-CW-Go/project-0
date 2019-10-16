# Project 0 Employee Reimbursement App
## Garner Deng
This is a basic Employee Reimbursement App.

You can make employee accounts with a username, password, first name and last name. Employees can submit and view their reimbursement tickets. In their ticket, they will
indicate an amount they are requesting as well as the reason.

You can also log in as a manager account to view all tickets, view only pending tickets, view all employees, and approve/deny reimbursement tickets. Manager accounts
require a 'master' password to create, because this function should not be accessible to everyone. 

All data will persist to a database. Requires docker and postgres. 

## Example: Employee Reimbursement App
- [x] Employees should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
    - [x] Submit reimbursement ticket
    - [x] View pending, approved, and denied reimbursements
- [x] Managers should be able to:
    - [x] View employee reimbursements
    - [x] Approve/deny open reimbursement requests
- [x] All accounts and reimbursements should persist to files or a database
- [x] Usable from command line args, interactive text menus, or through HTTP
- [x] Basic validation and testing

# User Stories
- [x] List
- [x] Each
- [x] User
- [x] Story
    1. Started project as a Banking app. Initially wanted to use structs to contain information, but ran into issues of linking structs to databases. Implementing functions for the structs would have been easy, but having all information persist to a database made the idea impractical.
    2. Changed project to Employee Reimbursement app. Same basic idea that requires the same core concepts to create, but seemed like a bit less coding.
    3. Scrapped structs idea. Decided to have ALL information stored as rows in database tables. The pokemon demo code really helped.
    4. Ran into serious issues pulling rows from the database tables. Eventually discovered that it may have been caused by having multiple tables and multiple insert row statements in the same init.sql file. Once I removed all insert row statements from the original init.sql file, I could pull rows from the database again. This problem was only discovered once I started adding tables to the database, so early on when there was only one database, everything worked fine.
    5. Changed fields in te init.sql file. Everything stopped working. No information could be found online because it was also difficult to identify the problem in the first place. Was getting 'nil pointer' errors. By sheer luck, realized that whenever changes were made to the init.sql file, the entire docker container had to be reinitialized. Ran lines like 
        
        To stop an old obsolete database:
                docker container ls     
                docker stop (container name)
        
        To create the database:
                cd db
                docker build -t (newContainername) .
                docker run --name (newContainername) -d -p 5432:5432 (newContainername)
                docker start (newContainername)
        Then to run the program:       
                cd ..
                go run main.go

    Had to create a new container every time I changed something in the init.sql file. This was probably the most difficult issue to solve.
    6. Needed to assign every row in every table a unique id number. Discovered the use of serial numbers that auto incremented themselves. Made implementing some functions much easier. 
    7. I had to run these lines of code every time I wanted to interact with the database tables:
        
        datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		    "localhost", 5432, "postgres", "postgres", "postgres")
	    db, err := sql.Open("postgres", datasource)
	    defer db.Close()
	    if err != nil {
		    panic(err)
	    }

        So I tried to clean up my code by turning this into a function that I could call, OPEN(). This broke the database. I eventually found out that it was because I included "defer db.Close()" inside of the function when I should have kept it outside; having it within the function meant the database opened and closed immediately. 
    8. Basically finished all the basic necessities of the program. Just needed to clean it up and make it look nice. Adding in extra lines of code that format the text in the terminal to make it look nicer and read easier. 

# Instructions
## Insert environment, build, and execution documentation here:
        To stop an old obsolete database:
                docker container ls     
                docker stop (container name)
        
        To create the database:
                cd db
                docker build -t (newContainername) .
                docker run --name (newContainername) -d -p 5432:5432 (newContainername)
                docker start (newContainername)
        Then to run the program:       
                cd ..
                go run main.go