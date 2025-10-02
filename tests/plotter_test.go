package tests

import (
	"os"
	"path/filepath"
	"testing"
	"time"
	"os/exec"
)

func TestPlotterCreatesFile(t *testing.T) {
	plotsDir := "plots"
	os.MkdirAll(plotsDir, 0755)

	// Remove old plots to start fresh
	files, _ := os.ReadDir(plotsDir)
	for _, f := range files {
		os.Remove(filepath.Join(plotsDir, f.Name()))
	}

	// Run the plotter via "go run"
	cmd := exec.Command("go", "run", "./cmd/plotter", "-size", "1", "-out", plotsDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Plotter execution failed: %v", err)
	}

	// Wait briefly in case of delayed write
	time.Sleep(500 * time.Millisecond)

	// Check if a new plot file exists
	newFiles, _ := os.ReadDir(plotsDir)
	if len(newFiles) == 0 {
		t.Fatalf("No plot file was created in %s", plotsDir)
	}

	// Check size > 0
	info, _ := os.Stat(filepath.Join(plotsDir, newFiles[0].Name()))
	if info.Size() == 0 {
		t.Fatalf("Plot file was created but size is zero")
	}
}

