create table machine (
    index integer,
    name varchar,
    stock integer,
    brand varchar
);

create table drinklist (
    index integer,
    name varchar unique,
    brand varchar,
    prob float
);