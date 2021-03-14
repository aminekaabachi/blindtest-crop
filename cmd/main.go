package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	core "github.com/aminekaabachi/blindtest-crop/core"
	"github.com/h2non/filetype"
)

var (
	seconds    = flag.Int("seconds", 30, "Number of seconds to crop.")
	inputPath  = flag.String("input-path", "", "MP3 songs input folder path.")
	outputPath = flag.String("output-path", "./output", "MP3 croped songs output path.")
)

func validateInput() (fs.FileInfo, []fs.FileInfo) {
	input, err := os.Stat(*inputPath)

	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The provided input path '%s' is not valid or does not exist.\n", *inputPath)
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(*inputPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return input, files
}

func validateOutput() fs.FileInfo {
	output, err := os.Stat(*outputPath)
	if os.IsNotExist(err) {
		log.Printf("Output path '%s' doesn't exists.\n", *outputPath)
		os.Mkdir(*outputPath, os.ModePerm)
		log.Printf("Created the following output path: '%s'.\n", *outputPath)
		output, _ = os.Stat(*outputPath)
	}

	return output
}

func Main() {
	flag.Parse()

	if *inputPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	input, files := validateInput()

	log.Printf("Started cropping all MP3 inside the path '%s'.\n", input.Name())

	output := validateOutput()

	var wg sync.WaitGroup

	for _, f := range files {
		path := filepath.Join(*inputPath, f.Name())
		buf, _ := ioutil.ReadFile(path)
		if filetype.Is(buf, "mp3") {
			wg.Add(1)
			go core.Crop(&wg, path, output.Name())
			wg.Wait()
		}
	}
}
