package main

import (
	"fmt"
	"log"

	"github.com/petrabarus/github-action-committer-coverage-stats/internal"
	"github.com/petrabarus/github-action-committer-coverage-stats/internal/coverage"
)

func main() {
	// Load Config
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("%+v\n", config)

	// Load Coverage File
	coverage, err := coverage.LoadCoverageFiles(config.CoverageFiles)
	if err != nil {
		log.Fatalf("Failed to load coverage file: %v", err)
	}
	fmt.Printf("%+v\n", coverage)

	// Load Git Repository
	repository, err := internal.LoadGitRepository(config.Workspace)
	if err != nil {
		log.Fatalf("Failed to load git repository: %v", err)
	}
	fmt.Printf("%+v\n", repository)

	// Calculate Coverage Stats
	stats, err := internal.CalculateCoverageStats(coverage, repository)
	if err != nil {
		log.Fatalf("Failed to calculate coverage stats: %v", err)
	}

	// Output Coverage Stats
	outputConfig := internal.StatsConfig{}
	err = internal.OutputCoverageStats(stats, &outputConfig)
	if err != nil {
		log.Fatalf("Failed to output coverage stats: %v", err)
	}
}
