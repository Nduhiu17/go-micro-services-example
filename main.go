package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w,"Ooops",http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Opps"))
			return 
		}
		fmt.Fprintf(w,"Hello %s",d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!")
	})

	http.ListenAndServe(":9090", nil)
}
