package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
// GET requests wait 5 seconds then respond

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Printf("%v", r.Header)
		fmt.Fprintf(w, "Hello from the slow server")
	})

	http.ListenAndServe(":8080", nil)


}