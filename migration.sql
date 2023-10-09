-- First create sequence 

      CREATE SEQUENCE id_increment START 1 INCREMENT 1;

-- Create table

      CREATE TABLE IF NOT EXISTS accounts (id int DEFAULT nextval('id_increment') primary key , name string unique not null, created_at TIMESTAMP not null, updated_at TIMESTAMP not null);