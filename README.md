
# http+cockroachdb

Simple user management go application using http methods with cockroachdb.


## Run Locally

Clone the project

```bash
  git clone https://github.com/KaranLathiya/http-cockroachdb.git
```

Install Go version 1.21

```bash
  go install 1.21 
```
Install Cockroachdb on local machine

Read the migration.sql  



Install dependencies

```bash
  go mod tidy 
```

Start the server

```bash
  go run main.go
```


## Deployment

To deploy this project run

```bash
  go run main.go
```


## Routing

To add new user details  --POST

    http://localhost:8080/user/add
To see all user details --GET

    http://localhost:8080/alluserdetails
To see user details with specific ID --GET

    http://localhost:8080/user/details
To delete user details with specific ID --DELETE

    http://localhost:8080/user/delete
To update user details with specific ID --PUT

    http://localhost:8080/user/update
## Table Details


| column_name | data_type | is_nullable |              column_default              | generation_expression |              indices              | is_hidden |
|--------------|-----------|-------------|------------------------------------------|-----------------------|-----------------------------------|------------|
|  id          | INT8      |      f      | nextval('public.id_increment'::REGCLASS) |                       | {accounts_name_key,accounts_pkey} |     f  |
|  name        | STRING    |      f      | NULL                                     |                       | {accounts_name_key,accounts_pkey} |     f  |
|  created_at  | TIMESTAMP |      f      | NULL                                     |                       | {accounts_pkey}                   |     f  |
|  updated_at  | TIMESTAMP |      f      | NULL                                     |                       | {accounts_pkey}                   |     f  |



|Sequence name|Create Statement |
|--------------|----------------------------------------------------- |
| id_increment	|CREATE SEQUENCE public.id_increment MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT 1 START 1 |