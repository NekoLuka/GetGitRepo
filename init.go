package main

import (
	"fmt"
	"log"
	"os"
)

func Init() {
	envs := []string{"PORT", "GITURL"}

	for _, v := range envs {
		if os.Getenv(v) == "" {
			log.Fatal(fmt.Sprintf("%s not set but required", v))
		}
	}

	optEnvs := map[string]string{"FETCH_INTERVAL": "300"}

	for k, v := range optEnvs {
		if os.Getenv(k) == "" {
			if err := os.Setenv(k, v); err != nil {
				log.Fatal(err)
			}
		}
	}
}
