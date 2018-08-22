package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func getusers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	users := [2]User{User{"First", 33}, User{"Second", 25}}

	encoder.Encode(users)
}

func main() {

	http.HandleFunc("/", greet)
	http.HandleFunc("/users", getusers)
	http.ListenAndServe(":8080", nil)
}
