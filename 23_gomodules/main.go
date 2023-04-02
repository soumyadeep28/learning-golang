package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to the golang portal ")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//http.ListenAndServe(":4000 ", r)
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {

	fmt.Println("hey there mod users")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to golang request call </h1>"))

}

/*
now in go mod file you can check that gomod files have the files with all the files but still indirect is used so you can run
"go mod tidy" command to reverify that
also to verify with the gir has run "go mod verify   "
to cehck all dependencies  you can run
go list
go list all
go list -m all
*/
