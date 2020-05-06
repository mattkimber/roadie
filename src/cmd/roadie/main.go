package main

import (
	"dto"
	"flag"
	"fmt"
	"log"
	"os"
)

type Flags struct {
	Version int
}

var flags Flags

const versionFile string = ".roadie_version"

func init() {
	flag.IntVar(&flags.Version, "version", -1, "version of the output file (for use with build pipelines")
	flag.IntVar(&flags.Version, "v", -1, "shorthand for -version")
}

func main() {
	flag.Parse()

	checkVersion()

	for _, a := range flag.Args() {
		processSetFile(a)
	}
}

func checkVersion() {
	if flags.Version == -1 {
		if _, err := os.Stat(versionFile); err != nil {
			writeVersion(1)
		} else {
			flags.Version = readVersion() + 1
			writeVersion(flags.Version)
		}
	}
}

func readVersion() (v int) {
	f, err := os.Open(versionFile)
	defer f.Close()
	if err != nil {
		log.Fatalf("could not open version file %s: %v", versionFile, err)
	}
	if _, err := fmt.Fscanf(f, "%d", &v); err != nil {
		log.Fatalf("could read from open version file %s: %v", versionFile, err)
	}
	return
}

func writeVersion(version int) {
	f, err := os.Create(versionFile)
	defer f.Close()
	if err != nil {
		log.Fatalf("could not create version file %s: %v", versionFile, err)
	}

	fmt.Fprintf(f, "%d", version)
}

func processSetFile(filename string) {
	set, err := dto.FromFile(filename, flags.Version)
	if err != nil {
		log.Printf("error loading %s: %v", filename, err)
		return
	}

	err = set.Create()
	if err != nil {
		log.Printf("error saving %s: %v", filename, err)
	}
}
