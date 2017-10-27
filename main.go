package main

import (
	"github.com/naoina/denco"
	"log"
	"net/http"
	"fmt"
	"github.com/beevik/etree"
)

func main() {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{
		mux.GET("/ReportingTool", Index),
		mux.GET("/ReportingTool/:style", Style),
	})
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func Index(w http.ResponseWriter, r *http.Request, params denco.Params) {
	fmt.Fprintf(w, "Welcome to /ReportingTool!\n")
}

func Style(w http.ResponseWriter, r *http.Request, params denco.Params) {
	fmt.Fprintf(w, "您要创建 %s!\n", params.Get("style"))
}
