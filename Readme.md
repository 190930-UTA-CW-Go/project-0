# Welcome to the F.R.E.D. Banking System
(Friendship,Reliability,Experince,Drive)

Banking Application with Database persistance

# User Stories
- [x] Checkout Requirements.md for more!
- [x] Checkout Examples.md for more!
- [ ] open accounts for approved users
- [ ] Make the application run in the webbrowser
  - [ ] get rid of bugs
  - [ ] implement employee page
  - [ ] implement page for potential new costomers


# Instructions
wanna see how its coming along? go to http://ec2-18-218-210-74.us-east-2.compute.amazonaws.com 
please dont break anythign or hack me
only sign in function works right now


create the database docker image and run it. the code should still connect without any issues.
`go run main.go` in the directory where file is located or cloned directory
this will run the CLI application

if you would like to run the Website verion cd into the elweb/ directory and run PS. make sure server is running
`go run server.go` which will host the server on localhost:6060
you may then sign in with credentials found in the init.sql file only the user login works currently.
Employee server implementation coming soon

That's it!
