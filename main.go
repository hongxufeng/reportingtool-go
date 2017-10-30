package main

import (
	"github.com/naoina/denco"
	"log"
	"net/http"
	"fmt"
)

func main() {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{
		mux.POST("/Login",UserLogin),
		mux.GET("/ReportingTool/:style", Style),
	})
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func UserLogin(w http.ResponseWriter, r *http.Request, params denco.Params) {

}

func Style(w http.ResponseWriter, r *http.Request, params denco.Params) {
	fmt.Fprintf(w, "您要创建 %s!\n", params.Get("style"))
}
