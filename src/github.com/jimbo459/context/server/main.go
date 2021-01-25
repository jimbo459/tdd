package main

import (
	"context"
	"net/http"
)

func main(){
	http.HandleFunc(w http.ResponseWriter, r *http.Request)

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)




}