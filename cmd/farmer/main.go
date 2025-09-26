// cmd/farmer/main.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	challenge := flag.String("challenge", "default-seed", "Challenge string")
	plotsDir := flag.String("plots", "plots", "Directory containing plot files")
	flag.Parse()

	// List plot files
	files, err := os.ReadDir(*plotsDir)
	if err != nil {
		fmt.Printf("Error reading plots directory: %v\n", err)
		return
	}
	if len(files) == 0 {
		fmt.Println("No plots found. Run the plotter first.")
		return
	}

	fmt.Printf("Farming with challenge: %s\n", *challenge)

	// Hash the challenge
	challengeHash := sha256.Sum256([]byte(*challenge))

	found := false
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		plotPath := filepath.Join(*plotsDir, f.Name())
		file, err := os.Open(plotPath)
		if err != nil {
			fmt.Printf("Error opening plot: %v\n", err)
			continue
		}
		defer file.Close()

		// Pick offset based on challenge hash
		offset := int64(challengeHash[0]) * 1024 // 1KB step
		_, err = file.Seek(offset, io.SeekStart)
		if err != nil {
			fmt.Printf("Error seeking in plot %s: %v\n", plotPath, err)
			continue
		}

		// Read 1KB from offset
		buf := make([]byte, 1024)
		_, err = file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading plot %s: %v\n", plotPath, err)
			continue
		}

		// Combine challenge with data slice
		data := append(challengeHash[:], buf...)
		result := sha256.Sum256(data)
		resultHex := hex.EncodeToString(result[:])

		fmt.Printf("Checked plot %s -> proof hash %s\n", f.Name(), resultHex[:16])

		// Simple "validity": check if hash starts with "00"
		if resultHex[:2] == "00" {
			fmt.Printf("Valid proof found in plot: %s\n", f.Name())
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No valid proof found for this challenge.")
	}
}

