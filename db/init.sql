create table employees (
    id serial primary key,
    fname varchar,
    lname varchar,
    username varchar unique,
    pass varchar,
    reimbursement varchar,
    currentstatus varchar
);

insert into employees values (default, 'terrell', 'green', 'terrellg', 'password','5000','pending');
insert into employees values(default, 'john', 'jay','jj123', 'password2', '10000','approved');