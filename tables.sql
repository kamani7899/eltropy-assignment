CREATE TABLE customers( id SERIAL PRIMARY KEY,  NAME  varchar(50) NOT NULL, AGE   INT  NOT NULL, ADDRESS  varchar(50),contact  varchar(50), created timestamp, active boolean, KYC varchar(30) not null);


create table accounts( id SERIAL PRIMARY KEY, accountType varchar(20), currentBalance real, status boolean, createdAt timestamp, Details varchar(50), customerid integer REFERENCES customers (id) );


CREATE TABLE employees( id SERIAL PRIMARY KEY, accountId integer REFERENCES accounts (id), NAME           varchar(50)    NOT NULL, AGE            INT     NOT NULL, ADDRESS  varchar(50),contact  varchar(50), createdAt timestamp, active boolean, salary real, admin boolean);


CREATE TABLE transactions( id SERIAL PRIMARY KEY, accountId integer REFERENCES accounts (id) , merchant varchar(50), amount real, transactionTime timestamp, employeeId integer REFERENCES employees (id), isSuccess boolean)