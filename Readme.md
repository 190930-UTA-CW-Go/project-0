# Bank application


## Nadine Gadjou

 CLI application that manage a bank.
##
It has twoo type of users (employees and customers).
The user can apply to create a new account, deposit or withdraw money. He can also create a joint account.
The employee can see all account infos, approuve or delete and application.
All data are saved into a postgres database image running in docker 

# User Stories
[] Customers should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
    - [x] Apply to open an account
    - [x] Apply for joint accounts with other customers
    - [x] Withdraw, deposit
[] Employees should be able to:
    - [X] View customer information
    - [X] Approve/deny open applications for accounts

# Instructions


To execute the program,
##
-install go , set  the PATH and the workspace folder. Then cd to the location that you want to put the program and clone the project.

-install docker: sudo apt install docker or sudo apt install docker.io
install postgres image 
-cd into the db folder where Dockerfile is running and run: 
-sudo docker build -t dbName .   then
-sudo docker run --name dbName -d -p 5432:5432 dbName
##

cd .. back to your main project 

Use go run *.go to run the program
