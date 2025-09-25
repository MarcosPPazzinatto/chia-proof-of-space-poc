// cmd/plotter/main.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	sizeMB := flag.Int("size", 10, "Size of the plot file in MB")
	outputDir := flag.String("out", "plots", "Directory to store plot files")
	flag.Parse()

	// Ensure output directory exists
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Generate filename with timestamp
	filename := fmt.Sprintf("plot-%d.dat", time.Now().Unix())
	filepath := filepath.Join(*outputDir, filename)

	// Open file for writing
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write pseudo-random data
	rand.Seed(time.Now().UnixNano())
	buffer := make([]byte, 1024*1024) // 1 MB buffer
	totalWritten := 0
	for totalWritten < (*sizeMB * 1024 * 1024) {
		rand.Read(buffer)
		n, err := file.Write(buffer)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return
		}
		totalWritten += n
	}

	// Compute SHA256 hash of file path and size (simple ID)
	idData := fmt.Sprintf("%s|%d", filepath, *sizeMB)
	hash := sha256.Sum256([]byte(idData))
	plotID := hex.EncodeToString(hash[:])

	fmt.Printf("Plot created: %s (%d MB)\n", filepath, *sizeMB)
	fmt.Printf("Plot ID: %s\n", plotID)
}

