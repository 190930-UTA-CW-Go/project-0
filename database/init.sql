create table users (
    usrID VARCHAR PRIMARY KEY UNIQUE,
    username varchar unique,
    passwd VARCHAR NOT NULL,
    fName VARCHAR NOT NULL,
    lName VARCHAR NOT NULL
);

insert into users values ('1','akhv','password','Artem','Khvan');
insert into users values ('2','jjohna','password','John','Johna');
insert into users values ('3','mmk','password','Michael','Moore');
insert into users values ('4','tcook','password','Tim','Cook');
insert into users values ('5','sjobs','password','Stev-o','Jobs');
insert into users values ('6','ckent','password','Clark','Kent');
insert into users values ('7','omg','password','Olly','Gunderson');
insert into users values ('8','imbored','password','Impending','Boredom');
insert into users values ('9','pparker','password','Peter','Parker');
insert into users values ('10','jonm','password','Jon','Martinez');

create table accounts (
    accNum VARCHAR NOT NULL UNIQUE,
    usrID VARCHAR REFERENCES users(usrID) NOT NULL,
    accType VARCHAR NOT NULL,
    accBal NUMERIC(10,2) NOT NULL ,
    PRIMARY KEY (accNum,usrID)
);

insert into accounts values ('1', '1', 'c', 1120.23);
insert into accounts values ('2', '2', 'c', 3000.23);
insert into accounts values ('3', '3', 'c', 4500.23);
insert into accounts values ('4', '4', 'c', 6500.23);
insert into accounts values ('5', '5', 'c', 300.23);
insert into accounts values ('6', '6', 'c', 530.23);
insert into accounts values ('7', '7', 'c', 9500.23);
insert into accounts values ('8', '8', 'c', 1320.23);
insert into accounts values ('9', '9', 'c', 15000.23);
insert into accounts values ('10', '10', 'c', 54000.23);
insert into accounts values ('11', '1', 's', 5060.23);
insert into accounts values ('12', '7', 's', 5040.23);
insert into accounts values ('13', '3', 's', 5000.23);

create table employees (
    empid VARCHAR PRIMARY KEY unique,
    username VARCHAR UNIQUE NOT NULL,
    pass VARCHAR NOT NULL,
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL
);

insert into employees VALUES ('1','emp1','password','Kerry','Reyes');
insert into employees VALUES ('2','emp2','password','James','Warren');
insert into employees VALUES ('3','emp3','password','John','Ferner');
insert into employees VALUES ('4','emp4','password','Maxwell','Shapiro');
insert into employees VALUES ('5','emp5','password','Lelouch','DeParte');

create table applications (
    username VARCHAR PRIMARY KEY unique,
    pass VARCHAR NOT NULL,
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL,
    actype VARCHAR NOT NULL,
    bal NUMERIC(10,2) NOT NULL
);

INSERT INTO applications VALUES ('sobas','password','Jon','Soba','c', 4321.50);
INSERT INTO applications VALUES ('stizzy','password','Sobek','Zaire','c', 5000.50);
INSERT INTO applications VALUES ('toolat3','password','Tophar','Laten','c', 4321.50);