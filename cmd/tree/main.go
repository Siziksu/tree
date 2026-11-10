package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	onlyDir := flag.Bool("d", false, "Show directories only")
	exclude := flag.String("e", "", "Comma-separated list of folders to exclude (e.g., .git,node_modules)")
	showHidden := flag.Bool("h", false, "Show hidden files and folders")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: tr [--d] [--e folder1,folder2,...] [depth]\n\nFlags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	excludedDirs := make(map[string]bool)
	if *exclude != "" {
		for dir := range strings.SplitSeq(*exclude, ",") {
			trimmed := strings.TrimSpace(dir)
			if trimmed != "" {
				excludedDirs[trimmed] = true
			}
		}
	}

	maxDepth := -1

	args := flag.Args()
	if len(args) > 0 {
		depth, err := strconv.Atoi(args[0])
		if err != nil || depth < 0 {
			fmt.Fprintf(os.Stderr, "Depth argument must be a no negative integer\n")
			flag.Usage()
			os.Exit(1)
		}
		maxDepth = depth
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error retrieving the current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(filepath.Base(currentDir))
	printTree(currentDir, "", 1, maxDepth, *onlyDir, *showHidden, excludedDirs)
}

func printTree(path string, prefix string, currentDepth int, maxDepth int, onlyDir bool, showHidden bool, excludedDirs map[string]bool) {
	if maxDepth != -1 && currentDepth > maxDepth {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if !showHidden && entry.Name()[0] == '.' {
			continue
		}
		if entry.IsDir() && excludedDirs[entry.Name()] {
			continue
		}
		if onlyDir && !entry.IsDir() {
			continue
		}
		filteredEntries = append(filteredEntries, entry)
	}

	for i, entry := range filteredEntries {
		isLast := i == len(filteredEntries)-1

		connector := "├── "
		if isLast {
			connector = "└── "
		}

		displayName := entry.Name()
		fmt.Printf("%s%s%s\n", prefix, connector, displayName)

		if entry.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			childPath := filepath.Join(path, entry.Name())
			printTree(childPath, newPrefix, currentDepth+1, maxDepth, onlyDir, showHidden, excludedDirs)
		}
	}
}
