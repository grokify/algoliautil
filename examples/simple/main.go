package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/grokify/algoliautil"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

const ENV_VAR = "ALGOLIA_APP_CREDENTIALS"

type Options struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
	Index   string `short:"i" long:"index" description:"Index" required:"true"`
	Add     string `short:"a" long:"add" description:"Add"`
	Query   string `short:"q" long:"query" description:"Query"`
}

type Contact struct {
	Name         string    `json:"name"`
	Value        int       `json:"value"`
	CreationTime time.Time `json:"createTime"`
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

	client, err := algoliautil.NewClientJSON([]byte(creds))
	if err != nil {
		log.Fatal(err)
	}

	index := client.InitIndex(opts.Index)

	addVar := strings.TrimSpace(opts.Add)

	if len(addVar) > 0 {
		rnd := rand.New(rand.NewSource(12345))
		contact := Contact{
			Name:         addVar,
			Value:        rnd.Int(),
			CreationTime: time.Now()}
		res, err := index.SaveObject(contact)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(res)
	}

	qryVar := strings.TrimSpace(opts.Query)

	if len(qryVar) > 0 {
		res, err := index.Search(opts.Query, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmtutil.PrintJSON(res)
		fmt.Printf("QRY_NUM_HITS [%v]\n", res.NbHits)
	}

	fmt.Println("DONE")
}
