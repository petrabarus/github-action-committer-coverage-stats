package coverage_test

import (
	"os"
	"path"
	"testing"

	"github.com/petrabarus/github-action-committer-coverage-stats/internal/coverage"
	"github.com/stretchr/testify/require"
)

func getFullPath(file string) string {
	// get current directory
	dir, _ := os.Getwd()
	file = path.Join(dir, file)
	return file
}

func TestLoadCoverageFromCoberturaXML(t *testing.T) {

	t.Run("test_case_001", func(t *testing.T) {
		// GIVEN
		file := "../../res/tests/cobertura-001.xml"
		file = getFullPath(file)

		// WHEN
		actual, err := coverage.LoadCoverageFromCoberturaXML(file)

		// THEN
		require.NoError(t, err)
		require.NotNil(t, actual)

		require.Equal(t, 4, len(actual.FileCoverage))

		firstFile := actual.FileCoverage[0]
		require.Equal(t, "src/github.rs", firstFile.FileName)
		require.Equal(t, 105, len(firstFile.LineCoverage))
	})
	t.Run("test_case_002", func(t *testing.T) {
		// GIVEN
		file := "../../res/tests/cobertura-002.xml"
		file = getFullPath(file)

		// WHEN
		actual, err := coverage.LoadCoverageFromCoberturaXML(file)

		// THEN
		require.NoError(t, err)
		require.NotNil(t, actual)
		require.Equal(t, 4, len(actual.FileCoverage))
		require.Equal(t, "Main.java", actual.FileCoverage[0].FileName)
		require.Equal(t, 11, len(actual.FileCoverage[0].LineCoverage))
	})
}
