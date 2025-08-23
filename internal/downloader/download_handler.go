package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Stats struct {
	Size      int64
	Timetaken time.Duration
}

func Download(link string, outputFile string) Stats {
	startTime := time.Now()
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Unable to create output file: %s\n", err)
		os.Exit(1)
	}
	defer out.Close()
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error making get request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Print("Bad response : $s\n", resp.StatusCode)
	}
	size, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Downloaded %d bytes to %s\n", size, outputFile)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	stats := Stats{
		Size:      size,
		Timetaken: elapsedTime,
	}
	return stats
}
