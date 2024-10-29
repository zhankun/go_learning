package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	router := httprouter.New()
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := w.Write([]byte("hello go http router"))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}

}
