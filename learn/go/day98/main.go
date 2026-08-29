package main

import (
	"fmt"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day98/internal/manifest"
)

func main() {
	fmt.Println("Day 98: deployment manifest (Kubernetes YAML demo)")
	fmt.Println(manifest.Summary())
	lines := strings.Split(strings.TrimSpace(manifest.DeploymentYAML), "\n")
	fmt.Printf("manifest lines: %d\n", len(lines))
	fmt.Println("first lines:")
	for _, line := range lines[:6] {
		fmt.Println(" ", line)
	}
	fmt.Println("  ...")
}
