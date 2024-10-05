package internal

type Coverage struct {
	FileCoverage []FileCoverage
}

type FileCoverage struct {
	FileName     string
	LineCoverage []LineCoverage
}

type LineCoverage struct {
	LineNumber int
	IsCovered  bool
}

func LoadCoverageFiles(paths []string) (*Coverage, error) {
	var coverage Coverage
	return &coverage, nil
}
