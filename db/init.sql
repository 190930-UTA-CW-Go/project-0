CREATE table newCustomer(
    id integer PRIMARY KEY,
    userName varchar,
	password varchar UNIQUE
);
insert into newCustomer values(1,'nadine', '1234');
insert into newCustomer values (2,'laure', '1235');

/* create table account(
    id integer primary key,
    custName  varchar,
	custNum     varchar unique,
	accountName  varchar ,
	accountNum   varchar unique,
	availableBal integer UNIQUE
); */


/* insert into account values (1, 'nadine',1234);
insert into account values (2, 'Ivysaur');
insert into account values (3, 'Venasaur'); */ 

