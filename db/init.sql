create table employee(
	name varchar(234) PRIMARY KEY,
	password varchar(222) NOT NULL

);


create table jointAccount(
	id integer primary key,
	name1 varchar(222) not null,
	name2 varchar(222) not null,
	accType varchar(222) not null,
	availableBal DECIMAL
);

CREATE TABLE account(
  	accountNum integer primary key,
    custName varchar (222) NOT NULL,
	age integer NOT NULL,
	accountType varchar(222),
	availableBal DECIMAL
); 

CREATE table newCustomer(
    name varchar(234) PRIMARY KEY,
	password varchar(222) NOT NULL
	
);





insert into employee values('nad','nad');

insert into employee values('nadine','nadine');


insert into newCustomer values('nadine','1234');




insert into jointAccount values(1,'laure','tamo','checking',100);
insert into jointAccount values (2, 'lauree','loui','saving',10);  

 


insert into account values (1234,'nadine', 23,'checking',100.33);
insert into account values (2345, 'bebe',30,'saving',200.33);
insert into account values (12345, 'tatm',60, 'checking',25.34);  

