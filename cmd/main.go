package cmd

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	seconds    = flag.Int("seconds", 30, "Number of seconds to crop.")
	inputPath  = flag.String("input-path", "", "MP3 songs input folder path.")
	outputPath = flag.String("output-path", "./output", "MP3 croped songs output path.")
)

func Main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()

	if *inputPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The provided input path is not valid or does not exist.\n")
		os.Exit(1)
	}
}
