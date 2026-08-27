package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputPath := filepath.Join("data", "day13", "sessions.log")
	outputPath := filepath.Join("day13", "summary.txt")

	data, err := os.ReadFile(inputPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file missing:", inputPath)
		}
		fmt.Println("read error:", err)
		return
	}

	fmt.Println("--- file content ---")
	fmt.Print(string(data))

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var builder strings.Builder
	writer := bufio.NewWriter(&builder)
	for _, line := range lines {
		fmt.Fprintln(writer, "session:", line)
	}
	writer.Flush()

	if err := os.WriteFile(outputPath, []byte(builder.String()), 0644); err != nil {
		fmt.Println("write error:", err)
		return
	}
	fmt.Println("wrote", outputPath)
}
