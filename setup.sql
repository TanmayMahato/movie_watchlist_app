
CREATE DATABASE IF NOT EXISTS movieapp;


USE movieapp;


-- Watchlist table
create table IF NOT EXISTS watchlist(
    mvid int primary key auto_increment,
    name varchar(100),
    gen varchar(20),
    cat varchar(20),
    exp int  );

-- Watched table

create table IF NOT EXISTS watched(
    mvid int, 
    name varchar(100), 
    gen varchar(20), 
    cat varchar(20), 
    rate int ,
    foreign key(mvid) references watchlist(mvid));

-- Trigger for inserting in watched from watchlist table on delete operation

CREATE TRIGGER IF NOT EXISTS before_movie_delete 
BEFORE DELETE ON watchlist 
FOR EACH ROW 
INSERT INTO watched (mvid, name, gen, cat) 
VALUES (OLD.mvid, OLD.name, OLD.gen, OLD.cat);
