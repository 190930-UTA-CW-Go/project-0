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
