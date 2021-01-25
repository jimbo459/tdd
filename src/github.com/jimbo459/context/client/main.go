package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
// Submit get request if takes longer than 3 seconds cancel the context
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%v", string(body))
}