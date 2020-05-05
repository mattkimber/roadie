package main

import (
	"dto"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	for _, a := range flag.Args() {
		processSetFile(a)
	}
}

func processSetFile(filename string) {
	set, err := dto.FromFile(filename, 1)
	if err != nil {
		log.Printf("error loading %s: %v", filename, err)
		return
	}

	err = set.Create()
	if err != nil {
		log.Printf("error saving %s: %v", filename, err)
	}
}
