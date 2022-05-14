package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("job is done")
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
		}
	})
	http.ListenAndServe("localhost:8088", nil)
}
