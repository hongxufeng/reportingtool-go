package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/ReportingTool", FormBuild)
	r.HandleFunc("/UserLogin", UserLogin)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

}

func FormBuild(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "您要创建 %s!\n")
}
