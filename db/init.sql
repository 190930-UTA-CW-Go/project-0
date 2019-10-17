CREATE TABLE user_accounts (
    uniqname varchar primary key,
    userfirst varchar,
    userlast varchar,
    password varchar);

CREATE TABLE bank_accounts (
    acctnumber varchar primary key,
    uniqname varchar,
    acctbalance decimal,
    accttype varchar);

CREATE TABLE joint_accounts (
    acctnumber varchar primary key,
    uniqname1 varchar,
    uniqname2 varchar,
    acctbalance decimal,
    acctname varchar);

CREATE TABLE employee_info (
    emp_number varchar,
    emp_first varchar,
    emp_last varchar,
    emp_password varchar,
    manager boolean);

insert into bank_accounts values('5464','cardib',54,'checking');
insert into bank_accounts values('7732', 'sprinkles', 99, 'savings');
insert into user_accounts values ('cardib', 'cardi', 'belcalis', 'bodakyellow');
insert into employee_info values ('1', 'Michael', 'Scott', 'password',false);
insert into user_accounts values ('schrutefarms','dwight', 'schrute','password');
insert into user_accounts values ('BEASLEY', 'pamela', 'beasley','password');
insert into user_accounts values ('bigtuna', 'James','Halpert', 'password');
insert into user_accounts values ('sprinkles','angela','schrute','password');
insert into user_accounts values ('smartypants', 'oscar', 'martinez','password');
insert into joint_accounts values ('5572', 'cardib', 'sprinkles', 74,'party fund');
insert into joint_accounts values ('3372', 'bigtuna', 'BEASLEY', 150, 'wedding');
