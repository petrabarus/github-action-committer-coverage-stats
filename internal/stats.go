package internal

import "github.com/petrabarus/github-action-committer-coverage-stats/internal/coverage"

type Stats struct {
}

func CalculateCoverageStats(coverage *coverage.Coverage, repository *Repository) (*Stats, error) {
	var stats Stats
	return &stats, nil
}
