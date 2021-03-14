package core

import (
	"log"
	"sync"
)

func Crop(wg *sync.WaitGroup, path string, output string) {
	defer wg.Done()
	log.Printf("%s cropping complete.\n", path)
}
