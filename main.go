package main

import (
	"fmt"
	"os"
	"time"
)

func progressBar(fileNum, totalLenth int) {
	barLength := 50                               // set the max length of the bar on the sceen
	perc_done := (fileNum + 1) * 100 / totalLenth // percentage done
	bars := perc_done * barLength / 100           // actual bars number

	var bar string = "["

	// write actual number of bars according to the current iteration
	for range bars {
		bar += "#"
	}

	fmt.Print("\033[u\033[K") // set cursor back to original location
	fmt.Printf("%v] %v/%v (%v%%)\n", bar, fileNum+1, totalLenth, perc_done)
	time.Sleep(1 * time.Millisecond) // simulate work
}

func main() {
	files, _ := os.ReadDir("files") // get the contents of the folder
	fmt.Print("\033[s")             // get current curson position
	// range over contents and call the progressBar() per each element
	for file := range files {
		progressBar(file, len(files))
	}
}
