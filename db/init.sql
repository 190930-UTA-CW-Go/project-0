create table customerLogin (
    userName varchar primary key,
    password varchar NOT NULL,
    fname varchar NOT NULL,
    lname varchar NOT NULL,
);
insert into customerLogin values ('Bulbasaury', 'pAsSwOrdd', 'Bulb', 'saur');
insert into customerLogin values ('adfh', 'password', 'Kdfa', 'miakhalifa');


create table customerAccounts(
    routingNumber serial primary key,
    balance float,
    userName varchar,
    userName2 varchar,
)
insert into customerAccounts values (123123, 500, 'Bulbasaury', 'adfh' );


create table employeeAccounts(
    adminLogin varchar primary key,
    password varchar NOT NULL,
)
insert into employeeAccounts values ('god', 'password');
