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
	SongId int `json:"song_id" query:"song_id" doc:"The internal identifier of the song"`
}

type SongSearchInput struct {
	Group       string `query:"group" doc:"The creator of the song"`
	Song        string `query:"song" doc:"The name of the song"`
	ReleaseDate string `query:"release_date" doc:"The release date of the song in format DD.MM.YYYY"`
	Text        string `query:"text" doc:"The lyrics of the song"`
	Link        string `query:"link" doc:"Link to the song"`
	SortBy      string `query:"sort_by" doc:"The column to sort by"`
	PageNum     int    `query:"page_num" doc:"The number of the fetched page"`
	PageLimit   int    `query:"page_limit" doc:"The maximum number of songs per page"`
}

type SongSearchOutput struct {
	Songs      []SongFull `json:"songs" doc:"The returned songs"`
	PageNum    int        `json:"page_num" doc:"The number of the returned page"`
	TotalPages int        `json:"total_pages" doc:"The number of pages in total"`
}

type SongTextFetchInput struct {
	SongId    int `query:"song_id" doc:"The id of the song"`
	PageNum   int `query:"page_num" doc:"The number of the fetched page"`
	PageLimit int `query:"page_limit" doc:"The maximum number of paragraphs per page"`
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
	}, func(ctx context.Context, input *struct{ Body SongCreateInput }) (*struct{ Body SongFull }, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-song",
		Method:      http.MethodPut,
		Path:        "/api/v1/song",
		Summary:     "Update a song",
		Description: "Replace the contents of a song with song_id.",
	}, func(ctx context.Context, input *struct{ Body SongFull }) (*struct{ Body SongFull }, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-song",
		Method:      http.MethodDelete,
		Path:        "/api/v1/song",
		Summary:     "Delete a song",
		Description: "Deletes a song with song_id.",
	}, func(ctx context.Context, input *struct{ Body SongDeleteInput }) (*struct{ Body SongFull }, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-song",
		Method:      http.MethodGet,
		Path:        "/api/v1/song",
		Summary:     "Get songs",
		Description: "Fetches songs, filtered and with pagination.",
	}, func(ctx context.Context, input *SongSearchInput) (*struct{ Body SongSearchOutput }, error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-lyrics",
		Method:      http.MethodGet,
		Path:        "/api/v1/lyrics",
		Summary:     "Get songs",
		Description: "Fetches lyric paragraphs, filtered and with pagination.",
	}, func(ctx context.Context, input *SongTextFetchInput) (*struct{ Body SongTextFetchOutput }, error) {
		return nil, nil
	})
}

func logMiddleware(l *log.Logger, ctx huma.Context, next func(huma.Context)) {
	l.Println(ctx.Operation())
	next(ctx)
}

func CreateServer(l *log.Logger) humacli.CLI {
	internalLogMiddleware := func(ctx huma.Context, next func(huma.Context)) {
		logMiddleware(l, ctx, next)
	}

	return humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		router := http.NewServeMux()
		config := huma.DefaultConfig("My API", "1.0.0")
		config.DocsPath = ""

		api := humago.New(router, config)
		api.UseMiddleware(internalLogMiddleware)

		addRoutes(api)
		router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
		<html lang="en">
		<head>
		  <meta charset="utf-8" />
		  <meta name="viewport" content="width=device-width, initial-scale=1" />
		  <meta name="description" content="SwaggerUI" />
		  <title>SwaggerUI</title>
		  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
		</head>
		<body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
		<script>
		  window.onload = () => {
		    window.ui = SwaggerUIBundle({
		      url: '/openapi.json',
		      dom_id: '#swagger-ui',
		    });
		  };
		</script>
		</body>
		</html>`))
		})

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			l.Println("[INFO] Starting server on port", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})
}
