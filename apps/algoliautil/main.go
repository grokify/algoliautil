package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grokify/algoliautil"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

const ENV_VAR = "ALGOLIA_APP_CREDENTIALS"

type Options struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
	Index   string `short:"i" long:"index" description:"Index"`
	Query   string `short:"q" long:"query" description:"Query"`
}

func main() {
	opts := Options{}
	if _, err := flags.Parse(&opts); err != nil {
		log.Fatal(err)
	}

	if err := config.LoadEnvPathsPrioritized(opts.EnvFile, os.Getenv(ENV_VAR)); err != nil {
		log.Fatal(err)
	}

	fmt.Println(os.Getenv(ENV_VAR) + "\n")

	creds := strings.TrimSpace(os.Getenv(ENV_VAR))
	if len(creds) == 0 {
		log.Fatal("No Algolia Credentials")
	}

	client, err := algoliautil.NewClientFromJSONSearchOrAdmin([]byte(creds))
	if err != nil {
		log.Fatal(err)
	}

	index := client.InitIndex(opts.Index)

	res, err := index.Search(opts.Query, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(res)
	fmt.Printf("NUM_HITS [%v]\n", res.NbHits)

	fmt.Println("DONE")
}
