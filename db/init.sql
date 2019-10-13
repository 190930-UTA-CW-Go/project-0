create table customer (
    userName varchar primary key,
    password varchar NOT NULL,
    name varchar NOT NULL,
    balance float64 NOT NULL, 

);

insert into customer values ('Bulbasaury', 'pAsSwOrdd', 'Bulb', 10);
insert into customer values ('adfh', 'password', 'Kdfa', 5000);