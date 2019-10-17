# The Banking App
## David Chang
A terminal interactive text menu banking app with the features:

# User Stories
- [x] Customers should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
    - [x] Apply to open an account
    - [x] Apply for joint accounts with other customers
    - [x] Withdraw, deposit, and transfer funds between accounts
- [x] Employees should be able to:
    - [x] View customer information
    - [x] Approve/deny open applications for accounts
- [x] All account and user information should persist to files or a database
- [x] Usable from command line args, interactive text menus, or through HTTP
- [x] Basic validation and testing

### Tools & APIs
- [x] Golang 1.10+
- [] Go Standard Library Packages
    - [x] `fmt`
    - [] `io`
    - [] `os`
    - [x] `flag`
    - [] `log`
    - [] `http`
    - [] `testing`

### Go language features
- [x] primitive data types
- [x] arrays/slices
- [] maps
- [] struct
- [] interface
- [x] functions/methods
- [] unit tests/benchmarks

### Functionality
- [x] CRUD - Create, Read, Update, Delete data
- [x] CLI - command-line args and/or event-driven text menus
- [] Execute OS commands
- [x] Login - Authentication & Authorization
- [x] Persisting State - Save/Load to file or DB
- [] HTTP - API endpoints accessible through browser endpoints, HTML/JS, and/or `curl` commands

# Instructions
To start database:
```bash
cd db
docker build -t project .
docker run --name mydb -d -p 5432:5432 project
```

To reset database:
```bash
docker stop mydb
docker rm mydb
docker rmi project
```

To check database:
```bash
docker ps -a
docker images -a
```

To start program:
```bash
go run *.go
go run *.go -hide
```