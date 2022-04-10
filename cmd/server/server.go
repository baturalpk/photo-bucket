package main

import (
	"log"

	"github.com/baturalpk/photo-bucket/api"
	"github.com/baturalpk/photo-bucket/clients/entclient"
	"github.com/baturalpk/photo-bucket/clients/s3client"
	"github.com/baturalpk/photo-bucket/config"
)

func main() {
	config.LoadConfigurations()

	if err := s3client.InitConnection(); err != nil {
		log.Fatalf("s3: init: error: %s\n", err)
	}

	if err := entclient.InitConnection(); err != nil {
		log.Fatalf("ent: init: error: %s\n", err)
	}

	if err := api.Serve(80); err != nil {
		entclient.CloseConnection()
		log.Fatalln(err)
		return
	}
}
