package main

import (
	"flag"
	"fmt"
	"net/url"
	"path"
	"sync"
	"time"

	"github.com/ahhossain/SuperGo/internal/downloader"
	"github.com/ahhossain/SuperGo/internal/stitcher"
)

func main() {
	dLink := flag.String("link", "", "Download Link")
	basePath := flag.String("path", "", "Download path")
	flag.Parse()
	dlLink := "https://releases.ubuntu.com/25.04/ubuntu-25.04-desktop-amd64.iso"
	parsedURL, err := url.Parse(dlLink)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	urlPath := parsedURL.Path
	outputFile := path.Base(urlPath)
	fmt.Println("Inferred filename from URL path:", outputFile)
	contentLength := downloader.GetHead(*dLink)
	outputPath := *basePath + outputFile
	if contentLength > 0 {
		fmt.Printf("Content length is: %d bytes\n", contentLength)
	} else {
		fmt.Println("Could not determine file size. Server may not support Content-Length header.")
	}
	chunks := createChunks(contentLength, downloader.ChunkSize, outputPath)
	chunkChan := make(chan downloader.Chunk, len(chunks))
	var wg sync.WaitGroup
	numWorkers := 12
	startTime := time.Now()
	for range numWorkers {
		wg.Add(1) // Increment the WaitGroup counter for each worker
		go func() {
			defer wg.Done()
			for chunk := range chunkChan {
				_ = downloader.GetChunk(chunk, *dLink, &wg)
				fmt.Printf("Chunk %d: Start=%d, End=%d, FileName=%s\n", chunk.ChunkNumber, chunk.StartChunk, chunk.EndChunk, chunk.ChunkName)
			}
		}()
	}
	for i := 1; i <= len(chunks); i++ {
		chunkChan <- chunks[i-1]
	}
	fmt.Println(len(chunks))
	close(chunkChan)
	wg.Wait()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Time Taken for full download %s\n", elapsedTime)
	fmt.Println("All chunks have been downloaded.")
	files := stitcher.Stitch(*basePath, outputFile, chunks)
	stitcher.DeleteChunks(files)

}

func createChunks(contentLength int64, chunkSize int64, outputPath string) []downloader.Chunk {
	var chunks []downloader.Chunk
	totalChunks := contentLength / chunkSize
	var i int64
	for i = 1; i <= totalChunks; i++ {
		tempFileName := fmt.Sprintf("%s.%d", outputPath, i)
		chunk := downloader.Chunk{
			ChunkNumber: i,
			ChunkName:   tempFileName,
			StartChunk:  chunkSize * (i - 1),
			EndChunk:    (chunkSize * i) - 1,
		}
		chunks = append(chunks, chunk)
	}
	tempFileName := fmt.Sprintf("%s.%d", outputPath, totalChunks+1)
	LastChunk := downloader.Chunk{
		ChunkNumber: i,
		ChunkName:   tempFileName,
		StartChunk:  chunkSize * (i - 1),
		EndChunk:    contentLength,
	}
	chunks = append(chunks, LastChunk)
	return chunks

}
