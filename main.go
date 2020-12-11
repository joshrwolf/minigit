package main

import (
	"github.com/sosedoff/gitkit"
	"log"
	"net/http"
	"os"
)

func main() {
	// Defaults from envs
	repoDir := os.Getenv("REPODIR")
	if repoDir == "" {
		repoDir = "/etc/minigit/repos/"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9696"
	}

	service := gitkit.New(gitkit.Config{
		Dir: repoDir,
		AutoCreate: true,
		AutoHooks: true,

	})

	if err := service.Setup(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", service)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}