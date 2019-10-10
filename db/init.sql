create database accounts;
\c accounts
create table users (
	username text primary key,
	password text not null
);
