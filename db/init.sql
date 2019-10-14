CREATE TABLE account(
  	accountNum varchar(200) PRIMARY KEY,
    custName varchar (222) NOT NULL,
	accountType varchar(222),
	availableBal DECIMAL
); 

CREATE table newCustomer(
    name varchar(234) PRIMARY KEY,
	password varchar(222) NOT NULL
	
);

/* insert into newCustomer values('nadine','1234');
insert into newCustomer values('laure','345');
insert into newCustomer values ('lauree','3456');  */

 


insert into account values ('1234','nadine','checking',100.33);
insert into account values ('2345', 'Ivysaur','saving',200.33);
insert into account values ('12345', 'Venasaur', 'checking',25.34);  

