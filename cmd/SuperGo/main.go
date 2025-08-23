package main

import (
	"fmt"

	"github.com/ahhossain/SuperGo/internal/downloader"
)

func main() {
	dlLink := "https://httpbin.org/bytes/1000000"
	//outputFile := "100MB.bin"
	contentLength := downloader.GetHead(dlLink)
	if contentLength > 0 {
		fmt.Printf("Content length is: %d bytes\n", contentLength)
	} else {
		fmt.Println("Could not determine file size. Server may not support Content-Length header.")
	}
	//stats := downloader.Download(dlLink, outputFile)
	//fmt.Printf("downloaded %d taking %s time\n", stats.Size, stats.Timetaken)
}
