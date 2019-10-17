# Banking App
## Khang Tran
Insert project description here.

# User Stories
10/8/2019 October 
-Created Employee and Customer Folder
-Created structs which will store user folders
- [x] List
- [] Each
- [] User
- [] Story

# Instructions
Insert environment, build, and execution documentation here.

To create the database:
```bash
cd db
docker build -t accountdb .
docker run --name accountdb -d -p 5432:5432 accountdb
```

Then use main.go to connect.

Open PostGressSQL using 
docker exec -it accountdb psql -U postgres