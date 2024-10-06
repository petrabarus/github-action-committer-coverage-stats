package internal

type Repository struct {
}

func LoadGitRepository(path string) (*Repository, error) {
	var repo Repository
	return &repo, nil
}
