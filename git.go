package main

import (
	"github.com/go-git/go-git/v5"
	"os"
	"strconv"
	"time"
)

type WikiGit struct {
	RepoLocation string
	Repo         *git.Repository
}

func (wg *WikiGit) Init() {
	logger := GetLogger()

	wg.RepoLocation = os.Getenv("GIT_REPO_LOCATION")

	if _, err := os.Stat(wg.RepoLocation); os.IsNotExist(err) {
		if err := os.MkdirAll(wg.RepoLocation, 0644); err != nil {
			logger.Error.Fatal(err)
		}

		repo, err := git.PlainClone(wg.RepoLocation, false, &git.CloneOptions{
			URL: os.Getenv("GITURL"),
		})
		if err != nil {
			logger.Error.Fatal(err)
		}

		wg.Repo = repo
		logger.Info.Println("Successfully cloned repo")
	} else {
		repo, err := git.PlainOpen(wg.RepoLocation)
		if err != nil {
			logger.Error.Fatal(err)
		}

		wg.Repo = repo
		logger.Info.Println("Successfully opened repo")
	}
}

func (wg *WikiGit) FetchAndPull() {
	logger := GetLogger()

	switch err := wg.Repo.Fetch(&git.FetchOptions{}); err {
	case nil:
	case git.NoErrAlreadyUpToDate:
		logger.Info.Println("No new data to be fetched")
		return
	default:
		logger.Error.Println(err)
		return
	}

	tree, err := wg.Repo.Worktree()
	if err != nil {
		logger.Error.Println(err)
		return
	}

	switch err = tree.Pull(&git.PullOptions{}); err {
	case nil:
		logger.Info.Println("Pull completed successfully")
	case git.NoErrAlreadyUpToDate:
		logger.Info.Println("No new data to pull")
	default:
		logger.Error.Println(err)
	}
}

func (wg *WikiGit) LoopFetchAndPull() {
	interval, _ := strconv.Atoi(os.Getenv("FETCH_INTERVAL"))
	for true {
		time.Sleep(time.Duration(interval) * time.Second)
		wg.FetchAndPull()
	}
}
