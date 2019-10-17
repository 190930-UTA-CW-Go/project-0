create table accounts (
    firstname varchar unique primary key,
    lastname varchar,
    password varchar,
    username varchar
);

create table customers (
    username varchar unique primary key,
    password varchar,
    firstname varchar,
    lastname varchar,
    balance FLOAT
);

create table employees (
    username varchar unique primary key,
    password varchar,
    firstname varchar,
    lastname varchar
);

insert into customers values ('DIO', 'ZAWARUDO', 'Dio','Brando',5000.75);
insert into customers values ('JOJO', 'ORAORA', 'Jotaro','Kujo',200.75);
insert into customers values ('Joker', 'Arsene', 'Ren','Amamiya',250.75);

insert into accounts values ('Tyler', 'Scott', 'Winnable','Tyler1');
insert into accounts values ('Billy', 'Batson','DC','SHAZAM');
insert into accounts values ('Shaquielle', 'O Neal','Dunkmaster','SHAQ');

insert into employees values ('George12','i<3pancakes','George','Conrad');
insert into employees values ('Ktran21','1234','Khang','Tran');
insert into employees values ('Brendo','i<3Tzuyu','Brendan','Xiong');