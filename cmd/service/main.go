package main

import (
	"context"
	"fmt"

	"github.com/serchemach/effective-mobile-test-task/api"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target api/ -package api --clean song_detail_scheme.yml
func main() {
	client, err := api.NewClient("http://localhost:8090")
	if err != nil {
		fmt.Println(err)
	}

	result, err := client.InfoGet(context.Background(), api.InfoGetParams{Group: "123", Song: "123"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
