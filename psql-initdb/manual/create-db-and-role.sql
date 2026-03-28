-- Required variables:
-- :db_name
-- :db_user
-- :db_pass

-- 1. Create role if it doesn't exist
-- (Native since PG 16)
CREATE USER :"db_user" WITH LOGIN PASSWORD :'db_pass';

-- 2. Create database if it doesn't exist 
CREATE DATABASE :"db_name" OWNER :"db_user";

-- 3. Grant privileges 
-- (Standard SQL, idempotent in this flow)
GRANT ALL PRIVILEGES ON DATABASE :"db_name" TO :"db_user";
