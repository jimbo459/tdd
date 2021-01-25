package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 4 * time.Second)

	defer cancel()

	select {
	case <- time.After(2 * time.Second):
		fmt.Println("Slept and awoke!")
	case <- ctx.Done():
		fmt.Println("Context cancelled")
	}

}
