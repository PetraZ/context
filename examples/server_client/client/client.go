package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	context, _ := context.WithTimeout(context.Background(), 1*time.Second)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8088", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req = req.WithContext(context)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
