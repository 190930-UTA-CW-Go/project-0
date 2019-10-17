create database accounts;
\c accounts
create table users (
	username text primary key,
	password text not null,
	status text not null
);
create table applications (
        username text primary key,
        firstname text not null,
        lastname text not null,
	address text not null,
	phone text not null
);
create table employees (
	username text primary key,
	password text not null
);
create table accountholders (
	username text primary key,
	accountnumber serial,
	firstname text not null,
	lastname text not null,
	address text not null,
	phone text not null,
	checking integer,
	savings integer
);
alter sequence accountholders_accountnumber_seq increment 1 restart with 1000000;
insert into employees (username, password) values ('employee1','password');
insert into employees (username, password) values ('employee2','password');
insert into employees (username, password) values ('employee3','password');
insert into employees (username, password) values ('employee4','password');
insert into employees (username, password) values ('employee5','password');
insert into users (username, password, status) values ('Bob','abc','pending');
insert into users (username, password, status) values ('Bill','abc','pending');
insert into applications (username,firstname,lastname,address,phone) values ('Bob','Robert','Smith','123 Main St. Dallas, TX 75001','469-123-4567');
insert into applications (username,firstname,lastname,address,phone) values ('Bill','William','Jones','45 Center St. Houston, TX 77001','281-123-4567');
