
create table employeeLogin (
    id SERIAL primary key,
    userName varchar unique,
    password varchar NOT NULL,
    fname varchar NOT NULL,
    lname varchar NOT NULL
);

create table pokemon (
    id integer primary key,
    name varchar unique
);

create table employeeAccounts(
    adminLogin varchar primary key,
    password varchar NOT NULL
);

create table tickets(
    ticketNum SERIAL primary key,
    userName varchar NOT NULL,
    fName varchar NOT NULL,
    lName varchar NOT NULL,
    reimburse float NOT NULL,
    reason varchar NOT NULL
);

