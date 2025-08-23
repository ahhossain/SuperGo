package main

import (
	"fmt"

	"github.com/ahhossain/SuperGo/internal/downloader"
)

func main() {
	dlLink := "https://sin-speed.hetzner.com/100MB.bin"
	outputFile := "100MB.bin"
	stats := downloader.Download(dlLink, outputFile)
	fmt.Printf("downloaded %d taking %s time\n", stats.Size, stats.Timetaken)
}
