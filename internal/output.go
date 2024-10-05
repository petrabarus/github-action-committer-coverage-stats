package internal

import (
	"fmt"
)

type StatsConfig struct {
	OutputFormat string
}

func OutputCoverageStats(stats *Stats, config *StatsConfig) error {
	fmt.Printf("%+v\n", stats)
	return nil
}
