package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const ChunkSize int64 = 4 * 1024 * 1024

type Stats struct {
	Size      int64
	Timetaken time.Duration
}

type Chunk struct {
	ChunkNumber int64
	ChunkName   string
	StartChunk  int64
	EndChunk    int64
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
	switch resp.StatusCode {
	case http.StatusPartialContent:
		contentRange := resp.Header.Get("Content-Range")
		if contentRange == "" {
			fmt.Println("Content-Range header missing.")
			return 0
		}
		parts := strings.Split(contentRange, "/")
		if len(parts) != 2 {
			fmt.Println("Invalid Content-Range header format.")
			return 0
		}
		totalSizeStr := parts[1]
		totalSize, err := strconv.ParseInt(totalSizeStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing total size from Content-Range: %v\n", err)
			return 0
		}
		return totalSize
	case http.StatusOK:
		return resp.ContentLength
	}
	fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
	return 0
}

func GetChunk(chunk Chunk, link string, wg *sync.WaitGroup) Stats {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}
	byteRange := fmt.Sprintf("bytes=%d-%d", chunk.StartChunk, chunk.EndChunk)
	req.Header.Set("Range", byteRange)
	client := &http.Client{}
	startTime := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making get request: %s\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		fmt.Printf("Bad response: %s\n", resp.Status)
	}
	out, err := os.Create(chunk.ChunkName)
	if err != nil {
		fmt.Printf("Unable to create output file: %s\n", err)
	}
	defer out.Close()
	size, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	stats := Stats{
		Size:      size,
		Timetaken: elapsedTime,
	}
	return stats
}
