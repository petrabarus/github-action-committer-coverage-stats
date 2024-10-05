package internal_test

import (
	"os"
	"testing"

	"github.com/petrabarus/github-action-committer-coverage-stats/internal"
	"github.com/stretchr/testify/require"
)

func TestLoadConfigFromEnv(t *testing.T) {

	t.Run("all_required_env_vars_set", func(t *testing.T) {
		// GIVEN
		os.Setenv("MIN_THRESHOLD", "80")
		os.Setenv("COVERAGE_FILES", "coverage.txt")
		os.Setenv("GITHUB_API_URL", "https://api.github.com")
		os.Setenv("GITHUB_TOKEN", "1234567890")
		os.Setenv("GITHUB_REF", "refs/heads/main")
		os.Setenv("GITHUB_REF_NAME", "main")
		os.Setenv("GITHUB_REPOSITORY", "petrabarus/github-action-committer-coverage-stats")
		os.Setenv("GITHUB_EVENT_NAME", "push")
		os.Setenv("GITHUB_HEAD_REF", "feature/test")

		// WHEN
		config, err := internal.LoadConfigFromEnv()

		// THEN
		require.NoError(t, err)
		require.Equal(t, 80, config.MinThreshold)
	})

	t.Run("missing_required_env_vars", func(t *testing.T) {
		// GIVEN
		os.Setenv("MIN_THRESHOLD", "80")
		os.Setenv("COVERAGE_FILES", "coverage.txt")
		os.Setenv("GITHUB_API_URL", "https://api.github.com")
		os.Setenv("GITHUB_TOKEN", "1234567890")
		os.Setenv("GITHUB_REF", "refs/heads/main")
		os.Setenv("GITHUB_REF_NAME", "main")
		os.Setenv("GITHUB_REPOSITORY", "petrabarus/github-action-committer-coverage-stats")
		os.Setenv("GITHUB_EVENT_NAME", "push")
		os.Setenv("GITHUB_HEAD_REF", "")

		// WHEN
		config, err := internal.LoadConfigFromEnv()

		// THEN
		require.Error(t, err)
		require.Nil(t, config)
	})
}
