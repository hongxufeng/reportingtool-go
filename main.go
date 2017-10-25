package reportingtool

import (
	"github.com/naoina/denco"
	"log"
	"net/http"
)

func main()  {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{
		mux.GET("/", Index),
	})
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}
