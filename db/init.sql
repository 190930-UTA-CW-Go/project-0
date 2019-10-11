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
	checking integer,
	savings integer
);
insert into employees (username, password) values ('employee1','password');
insert into employees (username, password) values ('employee2','password');
insert into employees (username, password) values ('employee3','password');
insert into employees (username, password) values ('employee4','password');
insert into employees (username, password) values ('employee5','password');
