create table customer (
    userID integer primary key,
    userName varchar unique NOT NULL,
    password varchar NOT NULL,
    firstName varchar NOT NULL,
    lastName varchar NOT NULL,
    balance int NOT NULL, 

);

insert into customer values (1, 'Bulbasaury', 'pAsSwOrdd', 'Bulba', 'Saur', 0);
