package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
		os.Exit(1)
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

func GetLength(link string) int64 {
	resp, err := http.Head(link)
	if err != nil {
		fmt.Printf("Error making get request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Print("Bad response : $s\n", resp.StatusCode)
		os.Exit(1)
	}
	contentLength := resp.ContentLength
	return contentLength
}

func GetHead(link string) int64 {
	var contentLength int64
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}
	req.Header.Set("Range", "bytes=0-0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making get request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusPartialContent || resp.StatusCode == http.StatusOK {
		contentLength, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			fmt.Printf("Error reading content-length from header: %v\n", err)
		}
		return contentLength
	}
	return contentLength
}
