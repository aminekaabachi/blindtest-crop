package core

import (
	"log"
	"os/exec"
	"strconv"
	"sync"
)

func Crop(wg *sync.WaitGroup, path string, output string, seconds int64) {
	defer wg.Done()
	cmd := exec.Command("ffmpeg", "-t", strconv.FormatInt(seconds, 10), "-i", path, "-acodec", "copy", "-y", output)
	cmd.Run()
	log.Printf("%s cropping complete.\n", path)
}
