# Banking App
## Khang Tran
--------------------------------BANKNG ZAWARUDO------------------------------------------
A Go CLI program. Update `Readme.md` with a name, a list of implemented and planned features, and instructions for running, building, testing, etc.

## Features
-Able to Create, Read, Update, and Delete from customers and employees tables
-Able to be run from the command line using user input
-Enforces the customer to login before he/she can utilize functionalities of the bank
-Combines PSQL with Golang to create a banking app, made possible by combining dockerfile with init.sql file along with other go files.
-
# User Stories
10/16/2019 October
Nearing project deadline, need to do a coding freeze to ensure that presentation goes smoothly despite unimplemented features.

10/12/2019 October
Decided to switch over to using a database to store all my information for my customer & employee

10/8/2019 October 
-Created Employee and Customer Folder
-Created structs which will store user folders

## Example: The Banking App
- [] Customers should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that username and password
    - [x] Apply to open an account
    - [] Apply for joint accounts with other customers
    - [x] Withdraw, deposit, and transfer funds between accounts
- [] Employees should be able to:
    - [x] View customer information
    - [] Approve/deny open applications for accounts
- [x] All account and user information should persist to files or a database
- [x] Usable from command line args, interactive text menus, or through HTTP

# Instructions
Insert environment, build, and execution documentation here.

To create the database:

```bash
cd into project-0 folder that contains sql file and docker file

cd go/src/github.com/project-0
docker build -t accountdb .
docker run --name accountdb -d -p 5432:5432 accountdb
```

Then use go run account.go main.go to connect.

Open PostGressSQL using 
docker exec -it accountdb psql -U postgres
Run \dt to see all tables inside of postgres
run select * from tablename;