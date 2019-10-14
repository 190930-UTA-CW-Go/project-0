create table customerLogin (
    userName varchar primary key,
    password varchar NOT NULL,
    fname varchar NOT NULL,
    lname varchar NOT NULL
);
insert into customerLogin values ('Bulbasaury', 'pAsSwOrdd', 'Bulb', 'saur');
insert into customerLogin values ('adfh', 'password', 'Kdfa', 'miakhalifa');



create table employeeAccounts(
    adminLogin varchar primary key,
    password varchar NOT NULL
);
insert into employeeAccounts values ('god', 'password');

create table pokemon (
    id integer primary key,
    name varchar unique
);

insert into pokemon values (1, 'Bulbasaur');
insert into pokemon values (2, 'Ivysaur');
insert into pokemon values (3, 'Venasaur');
insert into pokemon values (4, 'V1enasaur');
insert into pokemon values (5, 'Vena2saur');
insert into pokemon values (7, 'Vena2sa2ur');