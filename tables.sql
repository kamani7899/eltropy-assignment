CREATE TABLE customers( id SERIAL PRIMARY KEY,  NAME  varchar(50) NOT NULL, AGE   INT  NOT NULL, ADDRESS  varchar(50),contact  varchar(50), created timestamp, active boolean, KYC varchar(30) not null);


create table accounts( id SERIAL PRIMARY KEY, account_type varchar(20), current_balance real, status boolean, created_at timestamp, Details varchar(50), customer_id integer REFERENCES customers (id) );


CREATE TABLE employees( id SERIAL PRIMARY KEY, account_id integer REFERENCES accounts (id), NAME           varchar(50)    NOT NULL, AGE            INT     NOT NULL, ADDRESS  varchar(50),contact  varchar(50), created_at timestamp, active boolean, salary real, admin boolean);


CREATE TABLE transactions( id SERIAL PRIMARY KEY, account_id integer REFERENCES accounts (id) , merchant varchar(50), amount real, transaction_time timestamp, employee_id integer REFERENCES employees (id), is_success boolean)

create table users (username text primary key, password text, admin boolean, employee_id integer REFERENCES employees (id));