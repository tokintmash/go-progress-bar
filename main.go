package main

import (
	"fmt"
	"os"
)

func progressBar(filenum, totalLenth int) {
	length := 50
	perc_done := (filenum+1)*100/totalLenth
	bars := perc_done * length / 100

	var progressBar string = "["
	fmt.Print("\033[s")
	for i := 0; i < bars; i++ {
		progressBar += "#"
		fmt.Printf("%v",progressBar)
		// progressBar += fmt.Sprintf("%v %v %%", progressBar, perc_done)
	}
	fmt.Print("\033[u\033[K")
	// fmt.Printf("%v] %v %%\n", progressBar, (filenum+1)*100/totalLenth)
	// time.Sleep(1 * time.Millisecond)
}

func main() {
	files, _ := os.ReadDir("files")
	for file := range files {
		progressBar(file, len(files))
	}
	// for i := range files {
	// 	fmt.Print("\033[u\033[K")
	// 	progressBar += "#"
	// 	fmt.Printf("%v] %v %%\n", progressBar, (i+1)*100/len(files))
	// 	time.Sleep(1 * time.Millisecond)
	// }
}
