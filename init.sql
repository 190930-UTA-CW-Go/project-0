create table accounts (
    Firstname varchar unique primary key,
    Lastname varchar,
    Password varchar,
    Username varchar,
    Balance FLOAT
);

insert into accounts values ('Dio', 'Brando', 'DIO','ZAWARUDO',5000.75);
insert into accounts values ('Jotaro', 'Kujo', 'STAR Platinum','ORAORA',200.75);
insert into accounts values ('Ren', 'Amamiya', 'Joker','Arsene',250.75);