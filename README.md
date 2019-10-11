# The Banking App
## David Chang
Insert project description here.

# User Stories
- [x] List
- [] Each
- [] User
- [] Story

# Instructions
To start database:
cd db
docker build -t project .
docker run --name mydb -d -p 5432:5432 project

To reset database:
docker stop mydb
docker rm mydb
docker rmi project