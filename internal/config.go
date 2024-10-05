package internal

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	MinThreshold  int
	CoverageFiles []string
	BaseBranch    *string // optional
	FromTimestamp *string // optional
	ToTimestamp   *string // optional

	// Github Configurations
	GitHubApiUrl     string
	GitHubToken      string
	GitHubRef        string
	GitHubRefName    string
	GitHubRepository string
	GitHubEventName  string
	GitHubHeadRef    string
}

func LoadConfigFromEnv() (*Config, error) {
	var config Config

	err := loadBasicConfigFromEnv(&config)
	if err != nil {
		err = errors.New("failed to load basic config: " + err.Error())
		return nil, err
	}

	err = loadGithubConfigFromEnv(&config)
	if err != nil {
		err = errors.New("failed to load github config: " + err.Error())
		return nil, err
	}

	return &config, nil
}

func loadBasicConfigFromEnv(config *Config) error {

	minThreshold := os.Getenv("MIN_THRESHOLD")
	if minThreshold == "" {
		return errors.New("MIN_THRESHOLD is not set")
	}
	// parse minThreshold to int
	threshold, err := strconv.Atoi(minThreshold)
	if err != nil {
		return errors.New("MIN_THRESHOLD is not a number")
	}
	config.MinThreshold = threshold

	coverageFiles := os.Getenv("COVERAGE_FILES")
	if coverageFiles == "" {
		return errors.New("COVERAGE_FILES is not set")
	}
	config.CoverageFiles = strings.Split(coverageFiles, ",")

	baseBranch := os.Getenv("BASE_BRANCH")
	if baseBranch != "" {
		config.BaseBranch = &baseBranch
	}

	fromTimestamp := os.Getenv("FROM_TIMESTAMP")
	if fromTimestamp != "" {
		config.FromTimestamp = &fromTimestamp
	}

	toTimestamp := os.Getenv("TO_TIMESTAMP")
	if toTimestamp != "" {
		config.ToTimestamp = &toTimestamp
	}

	return nil
}

func loadGithubConfigFromEnv(config *Config) error {
	// GITHUB_API_URL
	config.GitHubApiUrl = os.Getenv("GITHUB_API_URL")
	if config.GitHubApiUrl == "" {
		return errors.New("GITHUB_API_URL is not set")
	}

	// GITHUB_TOKEN
	config.GitHubToken = os.Getenv("GITHUB_TOKEN")
	if config.GitHubToken == "" {
		return errors.New("GITHUB_TOKEN is not set")
	}

	// GITHUB_REF
	config.GitHubRef = os.Getenv("GITHUB_REF")
	if config.GitHubRef == "" {
		return errors.New("GITHUB_REF is not set")
	}

	// GITHUB_REF_NAME
	config.GitHubRefName = os.Getenv("GITHUB_REF_NAME")
	if config.GitHubRefName == "" {
		return errors.New("GITHUB_REF_NAME is not set")
	}

	// GITHUB_REPOSITORY
	config.GitHubRepository = os.Getenv("GITHUB_REPOSITORY")
	if config.GitHubRepository == "" {
		return errors.New("GITHUB_REPOSITORY is not set")
	}

	// GITHUB_EVENT_NAME
	config.GitHubEventName = os.Getenv("GITHUB_EVENT_NAME")
	if config.GitHubEventName == "" {
		return errors.New("GITHUB_EVENT_NAME is not set")
	}

	// GITHUB_HEAD_REF
	config.GitHubHeadRef = os.Getenv("GITHUB_HEAD_REF")
	if config.GitHubHeadRef == "" {
		return errors.New("GITHUB_HEAD_REF is not set")
	}

	return nil
}
