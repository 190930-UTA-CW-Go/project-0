create table accounts (
    firstname varchar unique primary key,
    lastname varchar,
    password varchar,
    username varchar,
    balance FLOAT
);

create table newcustomers(
    username varchar unique primary key,
    password varchar,
);

insert into accounts values ('Dio', 'Brando', 'DIO','ZAWARUDO',5000.75);
insert into accounts values ('Jotaro', 'Kujo', 'STAR Platinum','ORAORA',200.75);
insert into accounts values ('Ren', 'Amamiya', 'Joker','Arsene',250.75);

insert into newcustomers values ('Draven', 'Tyler1');
insert into newcustomers values ('SHAZAM', 'Billy');
insert into newcustomers values ('SHAQ', 'Lakers');