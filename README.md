
# http+cockroachdb

Simple user management go application using http methods with cockroachdb.


## Run Locally

Clone the project

```bash
  git clone https://github.com/KaranLathiya/http-cockroachdb.git
```

Go to the project directory

```bash
  cd my-project
```

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

    http://localhost:8080/add
To see all user details --GET

    http://localhost:8080/details
To see user details with specific ID --GET

    http://localhost:8080/idDetail
To delete user details with specific ID --DELETE

    http://localhost:8080/delete
To update user details with specific ID --PUT

    http://localhost:8080/update
## Table Details


| column_name | data_type | is_nullable |              column_default              | generation_expression |              indices              | is_hidden |
|--------------|-----------|-------------|------------------------------------------|-----------------------|-----------------------------------|------------|
|  id          | INT8      |      f      | nextval('public.id_increment'::REGCLASS) |                       | {accounts_name_key,accounts_pkey} |     f  |
|  name        | STRING    |      f      | NULL                                     |                       | {accounts_name_key,accounts_pkey} |     f  |
|  created_at  | TIMESTAMP |      f      | NULL                                     |                       | {accounts_pkey}                   |     f  |
|  updated_at  | TIMESTAMP |      f      | NULL                                     |                       | {accounts_pkey}                   |     f  |