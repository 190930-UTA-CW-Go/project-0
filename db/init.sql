create table bankaccount (
    name varchar unique,
    funds FLOAT,
    cfunds FLOAT,
    health INT,
    attack INT
);

insert into bankaccount (name, funds, cfunds, health, attack) values ('LokierFellheart', 123123, 123.32, 150, 20);
insert into bankaccount (name, funds, cfunds, health, attack) values ('BraveSirRobin', 42, 98, 20, 3);
insert into bankaccount (name, funds, cfunds, health, attack) values ('Bob', 32.23132, 653, 100, 10);
insert into bankaccount (name, funds, cfunds, health, attack) values ('TheDude', 32, 15, 8, 2);
insert into bankaccount (name, funds, cfunds, health, attack) values ('TheSpanishInquisition', 42, 19, 4, 9);
insert into bankaccount (name, funds, cfunds, health, attack) values ('ARandomRabbitThatWanderedIn', 14, 6, 18, 2);
insert into bankaccount (name, funds, cfunds, health, attack) values ('YourInnerDemons', 345, 64, 2, 20);