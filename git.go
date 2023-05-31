package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
)

type WikiGit struct {
	RepoLocation string
	Repo         *git.Repository
}

func (wg *WikiGit) Init() {
	wg.RepoLocation = os.Getenv("GIT_REPO_LOCATION")

	if _, err := os.Stat(wg.RepoLocation); os.IsNotExist(err) {
		if err := os.MkdirAll(wg.RepoLocation, 0644); err != nil {
			log.Fatal(err)
		}

		repo, err := git.PlainClone(wg.RepoLocation, false, &git.CloneOptions{
			URL: os.Getenv("GITURL"),
		})
		if err != nil {
			log.Fatal(err)
		}

		wg.Repo = repo
	} else {
		repo, err := git.PlainOpen(wg.RepoLocation)
		if err != nil {
			log.Fatal(err)
		}

		wg.Repo = repo
	}
}

func (wg *WikiGit) FetchAndPull() {
	err := wg.Repo.Fetch(nil)
	if !(err == nil || err == git.NoErrAlreadyUpToDate) {
		fmt.Print(err)
		return
	}

	tree, err := wg.Repo.Worktree()
	if err != nil {
		fmt.Print(err)
		return
	}

	err = tree.Pull(nil)
	if !(err == nil || err == git.NoErrAlreadyUpToDate) {
		fmt.Print(err)
		return
	}
	fmt.Print("Pull completed successfully")
}
