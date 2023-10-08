package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	// "github.com/rodaine/table"
	"strconv"

	_ "github.com/cockroachdb/cockroach-go/crdb"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error = nil

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func main() {
	connection_string := "postgresql://User:2a2dwrcFnaHmyS6I5mvE_A@solar-ape-6502.8nk.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err = sql.Open("postgres", connection_string)
	if err != nil {
		fmt.Println("Database Connection err", err)
		return
	}
	defer db.Close()

	// Create the "accounts" table.
	if _, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id int DEFAULT nextval('id_increment') primary key , name string unique not null, created_at TIMESTAMP not null, updated_at TIMESTAMP not null)"); err != nil {
		fmt.Println("Table creation err", err)
		return
	}
	//TO add new user details
	http.HandleFunc("/user/add", AddUserDetails) //POST
	//TO see all user details
	http.HandleFunc("/alluserdetails", AllUserDetails) //GET
	//TO see user details with specific ID
	http.HandleFunc("/user/details", UserDetailsById) //GET
	//TO delete user details with specific ID
	http.HandleFunc("/user/delete", DeleteUserDetails) //DELETE
	//TO update user details with specific ID
	http.HandleFunc("/user/update", UpdateUserDetails) //PUT
	fmt.Print("running")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func AddUserDetails(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user1 User
	json.Unmarshal(body, &user1)
	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid method")
		return
	}
	// name := r.FormValue("name")
	// if name == "" {
	// 	fmt.Fprintf(w, "Name value can't be empty.")
	// 	return
	// }
	// fmt.Println(name)
	if _, err := db.Exec(
		"INSERT INTO accounts (name,created_at,updated_at) VALUES ($1,$2,$2)", user1.Name, time.Now()); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "Data Successfully added")
}

func AllUserDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Invalid method")
		return
	}
	rows, err := db.Query("select  id, name, created_at, updated_at from accounts")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	var user1 User
	defer rows.Close()
	fmt.Fprint(w, "\nID\t\tNAME\t\t\t\tcreated_at\t\t\t\t\tupdated_at\n")
	for rows.Next() {
		if err := rows.Scan(&user1.Id, &user1.Name, &user1.Created_at, &user1.Updated_at); err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprintf(w, "%v\t|%-20v\t|%-20v\t|%-20v\n", user1.Id, user1.Name, user1.Created_at, user1.Updated_at)
	}
	// tbl := table.New("ID", "Name", "CREATED_AT","UPDATED_AT")
	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name,&created_at,&updated_at); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	tbl.AddRow(id,name,created_at,updated_at)
	// }
	// tbl.Print()
}

func UserDetailsById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Invalid method")
		return
	}
	var user1 User
	params := r.URL.Query()
	user1.Id, _ = strconv.Atoi(params.Get("id"))
	if user1.Id == 0 {
		fmt.Fprintf(w, "id value can't be empty.")
		return
	}
	errIfNoRows := db.QueryRow("select name,created_at,updated_at from accounts where id=$1", user1.Id).Scan(&user1.Name, &user1.Created_at, &user1.Updated_at)
	if errIfNoRows == nil {
		fmt.Fprintf(w, "id = %v \n name = %v \n created_at = %v \n updated_at = %v \n ", user1.Id, user1.Name, user1.Created_at, user1.Updated_at)
	} else {
		fmt.Fprintf(w, "%v id doesn't exist.", user1.Id)
	}
}

func DeleteUserDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprintf(w, "Invalid method")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user1 User
	json.Unmarshal(body, &user1)
	// fmt.Println(id)
	result, err := db.Exec(
		"delete from accounts where id= $1", user1.Id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintf(w, "RowsAffected Error %v", err)
		return
	}
	if RowsAffected == 1 {
		fmt.Fprintf(w, "Data of id %v is Successfully deleted", user1.Id)
	} else {
		fmt.Fprintf(w, "%v id doesn't exist.", user1.Id)
	}
}

func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		fmt.Fprintf(w, "Invalid method")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user1 User
	json.Unmarshal(body, &user1)
	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid method")
		return
	}
	// fmt.Println(id)
	result, err := db.Exec(
		"update accounts set name =$1,updated_at=$3 where id=$2", user1.Name, user1.Id, time.Now())
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintf(w, "RowsAffected Error %v", err)
		return
	}
	if RowsAffected == 1 {
		fmt.Fprintf(w, "Data of id %v is Successfully updated", user1.Id)
	} else {
		fmt.Fprintf(w, "%v id doesn't exist.", user1.Id)
	}

}
