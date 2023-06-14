package main

import (
	"log"
	"os"
)

func Init() {
	// Check for and initialize environment variables
	envs := []string{"GITURL"}

	for _, v := range envs {
		if os.Getenv(v) == "" {
			log.Fatalf("%s not set but required", v)
		}
	}

	optEnvs := map[string]string{
		"FETCH_INTERVAL":    "300",
		"PORT":              "5555",
		"GIT_REPO_LOCATION": "./repo",
		"LOG_FILE_LOCATION": "./RepoWiki.log",
		"LOG_LEVEL":         "1",
	}

	for k, v := range optEnvs {
		if os.Getenv(k) == "" {
			if err := os.Setenv(k, v); err != nil {
				log.Fatal(err)
			}
		}
	}
}
