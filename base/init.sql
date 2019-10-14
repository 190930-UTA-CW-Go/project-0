create table machine (
    index varchar unique,
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

create table services (
    account varchar unique,
    pass varchar,
    company varchar,
    firstname varchar,
    lastname varchar
);

insert into services values ('superd', 'SEC1verbose', 'Duda-Cola', 'D', 'Admin');
insert into services values ('supers', 'SEC2verbose', 'Salt-PhD',  'S', 'Admin');
insert into services values ('supert', 'SEC3verbose', 'TipsyCo',   'T', 'Admin');

insert into drinklist values (01, 'Duda-Cola',    'Duda-Cola', 5);
insert into drinklist values (02, 'Sprunk',       'Duda-Cola', 5);
insert into drinklist values (03, 'Diet Dude',    'Duda-Cola', 4);
insert into drinklist values (04, 'Dunce',        'Duda-Cola', 4);
insert into drinklist values (05, 'BBQ''s',       'Duda-Cola', 4);
insert into drinklist values (06, 'Fat-ah',       'Duda-Cola', 3);
insert into drinklist values (07, 'Electromaid',  'Duda-Cola', 3);
insert into drinklist values (08, 'Gold Snow',    'Duda-Cola', 3);
insert into drinklist values (09, 'Breaker',      'Duda-Cola', 2);
insert into drinklist values (10, 'Metal Pedal',  'Duda-Cola', 2);
insert into drinklist values (11, 'Ab Energy',    'Duda-Cola', 1);

insert into drinklist values (12, 'Salt-PhD',     'Salt-PhD', 5);
insert into drinklist values (13, 'Mexico-Wet',   'Salt-PhD', 5);
insert into drinklist values (14, 'Ph-1',         'Salt-PhD', 4);
insert into drinklist values (15, 'DDijon-Brown', 'Salt-PhD', 4);
insert into drinklist values (16, 'Q&A',          'Salt-PhD', 4);
insert into drinklist values (17, 'Squish',       'Salt-PhD', 3);
insert into drinklist values (18, 'SomeWatch',    'Salt-PhD', 3);
insert into drinklist values (19, 'Oh-Snap',      'Salt-PhD', 3);
insert into drinklist values (20, 'Fiji-Slap',    'Salt-PhD', 2);
insert into drinklist values (21, 'Poison',       'Salt-PhD', 2);
insert into drinklist values (22, 'Boohoo',       'Salt-PhD', 1);

insert into drinklist values (23, 'Tipsy',        'TipsyCo', 5);
insert into drinklist values (24, 'LowPool',      'TipsyCo', 5);
insert into drinklist values (25, 'ZeroTipsy',    'TipsyCo', 4);
insert into drinklist values (26, 'Blue-Gas',     'TipsyCo', 4);
insert into drinklist values (27, 'Mutt',         'TipsyCo', 4);
insert into drinklist values (28, 'Hot-Cabana',   'TipsyCo', 3);
insert into drinklist values (29, 'Geckoaide',    'TipsyCo', 3);
insert into drinklist values (30, 'Lipbalm',      'TipsyCo', 3);
insert into drinklist values (31, 'Warm',         'TipsyCo', 2);
insert into drinklist values (32, 'Hippie',       'TipsyCo', 2);
insert into drinklist values (33, 'Nude',         'TipsyCo', 1);