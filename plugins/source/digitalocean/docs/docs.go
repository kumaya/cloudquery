package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(provider.Provider(), "./docs", true)
	if err != nil {
		log.Fatalf("Failed to generate docs: %s", err)
	}
}
