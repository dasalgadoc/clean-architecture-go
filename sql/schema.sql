--- Script to initialize the database
CREATE SCHEMA test;
USE test;

create table courses (
id varchar(50),
name varchar(50),
PRIMARY KEY (id)
);
