/*This database will be used to track player playthroughs, 
showing either which ending of the game they got or where they died.*/

create table players (
    playername VARCHAR NOT NULL PRIMARY KEY,
    playerend VARCHAR NOT NULL
)