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
    email varchar references customer(email) on delete cascade,
    acc_type varchar,
    acc_balance float,
    acc_num serial unique
);

insert into employee values ('user', 'pass', 'David', 'Chang');