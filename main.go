package main

import (
	"fmt"
	"log"

	"github.com/petrabarus/github-action-committer-coverage-stats/internal"
)

func main() {
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("%+v\n", config)
}
