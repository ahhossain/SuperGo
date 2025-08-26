package stitcher

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ahhossain/SuperGo/internal/downloader"
)

func Stitch(basePath string, outputFile string, chunks []downloader.Chunk) []string {
	outputPath := basePath + outputFile
	var files []string
	fmt.Println(outputPath)
	destFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("failed to create destination file: %v", err)
	}
	fmt.Printf("Created output file: %s\n", outputPath)
	defer destFile.Close()
	for _, chunk := range chunks {
		srcFile, err := os.Open(chunk.ChunkName)
		if err != nil {
			fmt.Printf("failed to open source file 2: %v\n", err)
		}
		defer srcFile.Close()
		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			fmt.Printf("failed to copy file 1: %v\n", err)
		}
		fmt.Printf("Merged : %s\n", chunk.ChunkName)
		files = append(files, chunk.ChunkName)
	}
	return files
}

func DeleteChunks(files []string) {
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			log.Printf("failed to delete source file '%s': %v\n", file, err)
		} else {
			fmt.Printf("Successfully copied and deleted source file: %s\n", file)
		}
	}
}
