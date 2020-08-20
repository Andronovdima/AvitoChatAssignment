package main

import (
	apiserver "../../internal/app"
	"log"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
