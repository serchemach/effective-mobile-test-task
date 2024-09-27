package handling

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8099"`
}

type SongCreateInput struct {
	Group string `json:"group" doc:"The creator of the song"`
	Song  string `json:"song" doc:"The name of the song"`
}

type SongFull struct {
	SongCreateInput
	ReleaseDate string `json:"release_date" doc:"The release date of the song in format DD.MM.YYYY"`
	Text        string `json:"text" doc:"The lyrics of the song"`
	Link        string `json:"link" doc:"Link to the song"`
	SongId      int    `json:"song_id" doc:"The internal identifier of the song"`
}

type SongDeleteInput struct {
	SongId int `json:"song_id" doc:"The internal identifier of the song"`
}

type SongSearchInput struct {
	Group       string `json:"group" doc:"The creator of the song"`
	Song        string `json:"song" doc:"The name of the song"`
	ReleaseDate string `json:"release_date" doc:"The release date of the song in format DD.MM.YYYY"`
	Text        string `json:"text" doc:"The lyrics of the song"`
	Link        string `json:"link" doc:"Link to the song"`
	SortBy      string `json:"sort_by" doc:"The column to sort by"`
	PageNum     int    `json:"page_num" doc:"The number of the fetched page"`
	PageLimit   int    `json:"page_limit" doc:"The maximum number of songs per page"`
}

type SongSearchOutput struct {
	Songs      []SongFull `json:"songs" doc:"The returned songs"`
	PageNum    int        `json:"page_num" doc:"The number of the returned page"`
	TotalPages int        `json:"total_pages" doc:"The number of pages in total"`
}

type SongTextFetchInput struct {
	SongId    int `json:"song_id" doc:"The id of the song"`
	PageNum   int `json:"page_num" doc:"The number of the fetched page"`
	PageLimit int `json:"page_limit" doc:"The maximum number of paragraphs per page"`
}

type SongTextFetchOutput struct {
	Paragraphs []string `json:"paragraphs" doc:"The returned paragraphs"`
	PageNum    int      `json:"page_num" doc:"The number of the returned page"`
	TotalPages int      `json:"total_pages" doc:"The number of pages in total"`
}

func addRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-song",
		Method:      http.MethodPost,
		Path:        "/api/v1/song",
		Summary:     "Add a new song",
		Description: "Create a new song and fetch the details from a remote API.",
	}, func(ctx context.Context, input *SongCreateInput) (*SongFull, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-song",
		Method:      http.MethodPut,
		Path:        "/api/v1/song",
		Summary:     "Update a song",
		Description: "Replace the contents of a song with song_id.",
	}, func(ctx context.Context, input *SongFull) (*SongFull, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-song",
		Method:      http.MethodDelete,
		Path:        "/api/v1/song",
		Summary:     "Delete a song",
		Description: "Deletes a song with song_id.",
	}, func(ctx context.Context, input *SongDeleteInput) (*SongFull, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-song",
		Method:      http.MethodGet,
		Path:        "/api/v1/song",
		Summary:     "Get songs",
		Description: "Fetches songs, filtered and with pagination.",
	}, func(ctx context.Context, input *SongSearchInput) (*SongSearchOutput, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-lyrics",
		Method:      http.MethodGet,
		Path:        "/api/v1/lyrics",
		Summary:     "Get songs",
		Description: "Fetches lyric paragraphs, filtered and with pagination.",
	}, func(ctx context.Context, input *SongTextFetchInput) (*SongTextFetchOutput, error) {
		return nil, nil
	})
}

func CreateServer(l *log.Logger) humacli.CLI {
	return humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		router := http.NewServeMux()
		api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))

		addRoutes(api)

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			l.Println("[INFO] Starting server on port", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})
}
