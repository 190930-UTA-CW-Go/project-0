create table employeeLogin (
    id SERIAL primary key,
    userName varchar NOT NULL unique,
    password varchar NOT NULL,
    fname varchar NOT NULL,
    lname varchar NOT NULL
);
insert into employeeLogin values ('Bulbasaur', 'pAsSwOrdd', 'Bulb', 'Saur');
insert into employeeLogin values ('Ivysaur', 'password', 'Ivy', 'Saur');
insert into employeeLogin values ('BigDaddyKane', 'password', 'Antonio', 'Hardy');



create table employeeAccounts(
    adminLogin varchar primary key,
    password varchar NOT NULL
);
insert into employeeAccounts values ('god', 'password');

create table tickets(
    ticketNum SERIAL primary key,
    userName varchar NOT NULL,
    fName varchar NOT NULL,
    lName varchar NOT NULL,
    reimburse float NOT NULL,
    reason varchar NOT NULL,
);
insert into tickets values (1, 'BigDaddyKane', 'Antonio', 'Hardy', 9001, 'Gimme da monay');