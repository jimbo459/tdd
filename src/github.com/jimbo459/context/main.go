package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)

	defer cancel()

	sleepAndTalk(ctx, 3 * time.Second,"Slept and woke")

}

func sleepAndTalk(ctx context.Context,duration time.Duration, message string) {
	select {
	case <- time.After(duration):
		fmt.Println(message)
	case <- ctx.Done():
		fmt.Printf("%v",ctx.Err().Error())
	}

}