
package main

import(
	"net/http"
	"fmt"
	"encoding/json"
)

type User struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "Jack"},
	{ID: "2", Name: "Chitty"},
}

// controller

func getUsers(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request){
    var user User
	
	json.NewDecoder(r.Body).Decode(&user)

    users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello world")
}

func main(){
	http.HandleFunc("/", home)
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/users", getUsers)
	fmt.Println("server is tunnign at port")
	http.ListenAndServe(":8080", nil)
}