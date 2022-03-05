package main

import (
	"log"

	"github.com/baturalpk/photo-bucket/config"
)

func main() {
	config.LoadConfigurations()

	log.Println(config.Get())
}
