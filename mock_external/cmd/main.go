package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/serchemach/effective-mobile-test-task/api"
)

type mockExternal struct {
}

func (m *mockExternal) InfoGet(ctx context.Context, params api.InfoGetParams) (api.InfoGetRes, error) {
	log.Println(params)
	return &api.SongDetail{ReleaseDate: params.Group, Text: params.Song}, nil
}

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ../../api/ -package api --clean song_detail_scheme.yml
func main() {
	server, err := api.NewServer(&mockExternal{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SERVING MOCK")
	if err := http.ListenAndServe("0.0.0.0:8090", server); err != nil {
		log.Fatal(err)
	}
}
