package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/serchemach/effective-mobile-test-task/api"
	"github.com/serchemach/effective-mobile-test-task/service/db"
	"github.com/serchemach/effective-mobile-test-task/service/handling"
)

func envWithDefault(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func envWithError(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf("The %s variable is not present", key))
	}
	return val, nil
}

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ../../api/ -package api --clean ../../api/song_detail_scheme.yml
func main() {
	l := log.New(os.Stdout, "", 1)
	err := godotenv.Load()
	if err != nil {
		l.Fatalln(err)
	}

	postgresUser, err := envWithError("POSTGRES_USER")
	if err != nil {
		l.Fatalln(err)
	}

	postgresPass, err := envWithError("POSTGRES_PASSWORD")
	if err != nil {
		l.Fatalln(err)
	}

	postgresDb, err := envWithError("POSTGRES_DB")
	if err != nil {
		l.Fatalln(err)
	}

	postgresUrl := envWithDefault("POSTGRES_URL", "postgres")
	if err != nil {
		l.Fatalln(err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", postgresUser, postgresPass, postgresUrl, postgresDb)
	err = db.InitializeMigrations("file://migrations", connString)
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			l.Println("[INFO] Migrations were not necessary")
		} else {
			l.Fatalln(err)
		}
	} else {
		l.Println("[INFO] Successfully applied migrations")
	}

	mockPort := envWithDefault("MOCK_EXTERNAL_PORT", "8090")
	client, err := api.NewClient("http://app:" + mockPort)
	if err != nil {
		l.Fatalln(err)
	}
	l.Println("[INFO] Successfully initialized external api client")

	cli := handling.CreateServer(l)
	cli.Run()

	result, err := client.InfoGet(context.Background(), api.InfoGetParams{Group: "123", Song: "123"})
	if err != nil {
		l.Fatalln(err)
	}

	fmt.Println(result)
}
