package main

import (
	"log"

	"github.com/m507/subzy/src"
)

func main() {

	src.CheckFingerprints()

	settings := src.ParseConfiguration()

	if settings == nil {
		src.PrintHelp()
	}

	if err := src.Process(settings); err != nil {
		log.Fatalf("Error processing: %v\n", err)
	}

}
