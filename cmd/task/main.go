package main

import (
	"fmt"
	"github.com/Ablyamitov/task/internal/config"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	fmt.Println(conf)
}
