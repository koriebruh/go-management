CREATE  DATABASE  go_management;

USE go_management;

SHOW TABLES ;

desc admins;
select * from admins;

desc categories;
select * from categories;

desc suppliers;
select * from suppliers;

desc items;
select * from items;


CREATE DATABASE go_management_test;
USE go_management_test;
SHOW TABLES ;

select * from admins;
delete from admins;