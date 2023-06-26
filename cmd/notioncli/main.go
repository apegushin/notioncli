package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/apegushin/notioncli/pkg/config"
	"github.com/apegushin/notioncli/pkg/notion-sdk"
)

func main() {
	fmt.Printf("Starting main()\n")

	integrationName := flag.String("integration-name", "", "Symbolic name of a connection to Notion DB")
	integrationToken := flag.String("integration-token", "", "Token supplied for integration by Notion My-Integrations")
	integrationDatabaseId := flag.String("database-id", "", "ID of the database to which integration was applied")
	flag.Parse()

	// TODO: remove this and further extra logging
	fmt.Printf("Defined and parsed CLI args\n")

	if *integrationName != "" {
		err := config.NewConfig("notioncli.cfg").AddIntegrationRecord(*integrationName, *integrationToken, *integrationDatabaseId)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	notion := notionsdk.NewClient("notioncli.cfg")
	notion.Get()

	fmt.Printf("Success! All done!\n")
}
