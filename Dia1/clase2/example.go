package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	data, err := fetchData(ctx, "www.google.com")
	if err != nil {
		fmt.Println("request error: ", err)
		return
	}
	fmt.Println("request ok: ", data)
}

func fetchData(ctx context.Context, url string) (string, error) {
	fmt.Println("fetching url:", url)
	time.Sleep(time.Second * 3)
	return "data", nil
}
