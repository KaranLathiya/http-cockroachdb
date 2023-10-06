package main

import (
	"database/sql"
	"fmt"
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

func main(){
	dsn := "postgresql://User:2a2dwrcFnaHmyS6I5mvE_A@solar-ape-6502.8nk.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "accounts" table.
	if _, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id int DEFAULT nextval('id_increment') primary key ,name string unique not null,created_at TIMESTAMP not null ,updated_at TIMESTAMP not null)"); err != nil {
		log.Fatal(err)
	}
	//TO add new user details
	http.HandleFunc("/add", AddDetails)				//POST
	//TO see all user details
	http.HandleFunc("/details", AllDetails)			//GET
	//TO see user details with specific ID
	http.HandleFunc("/idDetail", IdDetails)			//GET
	//TO delete user details with specific ID
	http.HandleFunc("/delete", DeleteDetails)			//DELETE
	//TO update user details with specific ID
	http.HandleFunc("/update", UpdateDetails)			//PUT
	fmt.Print("running")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func AddDetails(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
	fmt.Fprintf(w, "Invalid method")
	} else{
		name := r.FormValue("name")
		if name==""{
			fmt.Fprintf(w, "Name value can't be empty.")
			return
		}
		// fmt.Println(name)
		if _, err := db.Exec(
			"INSERT INTO accounts (name,created_at,updated_at) VALUES ($1,$2,$3)",name,time.Now(),time.Now()); err != nil {
				fmt.Fprint(w, err)
		} else{
			fmt.Fprintf(w, "Data Successfully added")
		}
		}
}
func AllDetails(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		fmt.Fprintf(w, "Invalid method")
		} else{
			var id,name,created_at,updated_at string
			if rows,err := db.Query("select * from accounts"); err!=nil{
				fmt.Fprint(w, err)
			} else {
				defer rows.Close()
				fmt.Fprint(w,"\nID\t\tNAME\t\t\t\t\t\t\tcreated_at\t\t\t\t\tupdated_at\n")
				for rows.Next() {
					if err := rows.Scan(&id, &name,&created_at,&updated_at); err != nil {
						log.Fatal(err)
					}
					fmt.Fprintf(w,"%v\t|%-20v\t|%-20v\t|%-20v\n",id,name,created_at,updated_at)
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
			}	
		}

func IdDetails(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		fmt.Fprintf(w, "Invalid method")
		} else{
			params := r.URL.Query()
			id :=params.Get("id")
			if id==""{
				fmt.Fprintf(w, "id value can't be empty.")
				return
			}
			// fmt.Println(id)
			var name,created_at,updated_at string
			errIfNoRows := db.QueryRow("select name,created_at,updated_at from accounts where id=$1",id).Scan(&name,&created_at,&updated_at) 
			if errIfNoRows == nil{
				fmt.Fprintf(w, "id = %v \n name = %v \n created_at = %v \n updated_at = %v \n ",id,name,created_at,updated_at)
			} else{
				fmt.Fprintf(w, "%v id doesn't exist.",id)
			}
		}
			}

func DeleteDetails(w http.ResponseWriter, r *http.Request){
	if r.Method != "DELETE"{
		fmt.Fprintf(w, "Invalid method")
		} else{
			id := r.FormValue("id")
			if id==""{
				fmt.Fprintf(w, "id value can't be empty.")
				return
			}
			// fmt.Println(id)
			if result, err := db.Exec(
				"delete from accounts where id= $1",id); err != nil {
					fmt.Fprint(w, err)
			} else{
				RowsAffected, err := result.RowsAffected()
				if err != nil {
					fmt.Fprintf(w,"RowsAffected Error %v", err)
				}
				if RowsAffected==1{
				fmt.Fprintf(w, "Data of id %v is Successfully deleted",id)
				} else {
					fmt.Fprintf(w, "%v id doesn't exist.",id)
				}
			}
			}
}
func UpdateDetails(w http.ResponseWriter, r *http.Request){
	if r.Method != "PUT"{
		fmt.Fprintf(w, "Invalid method")
		} else{
			string_id := r.FormValue("id")
			name := r.FormValue("name")
			if (string_id=="" || name==""){
				fmt.Fprintf(w, "Null values not allowed")
				return
			}
			id,_:=strconv.Atoi(string_id)
			// fmt.Println(id)
			if result, err := db.Exec(
				"update accounts set name =$1,updated_at=$3 where id=$2",name,id,time.Now()); err != nil {
					fmt.Fprint(w, err)
			} else{
				RowsAffected, err := result.RowsAffected()
				if err != nil {
					fmt.Fprintf(w,"RowsAffected Error %v", err)
				}
				if RowsAffected==1{
				fmt.Fprintf(w, "Data of id %v is Successfully updated",id)
				} else {
					fmt.Fprintf(w, "%v id doesn't exist.",id)
				}
			}
			}
}
