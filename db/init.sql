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
    acc_num varchar,
    acc_balance float
);

insert into employee values ('user', 'pass', 'David', 'Chang');