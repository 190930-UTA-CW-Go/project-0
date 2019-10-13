create table customer (
    email varchar primary key,
    pass varchar, 
    first_name varchar,
    last_name varchar
);

create table employee (
    email varchar primary key,
    pass varchar,
    first_name varchar,
    last_name varchar
);

create table account (
    email varchar references customer(email),
    acc_type varchar,
    acc_balance float,
    acc_num serial
);

insert into employee values ('user', 'pass', 'David', 'Chang');