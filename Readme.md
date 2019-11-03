# Project 1 SSH 
## Garner Deng

All data will persist to a database. Requires docker and postgres. 

## 
- [] Usable from command line args, interactive text menus, or through HTTP
- [] Basic validation and testing

# User Stories

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
