package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	varIndex = make(map[string]bool)
	mu       sync.Mutex
)

func scanFile(path string, extensions []string) {
	if !isValidFile(path, extensions) {
		return
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return
	}

	re := regexp.MustCompile(`process\.env\.([a-zA-Z0-9_]+)`)
	matches := re.FindAllStringSubmatch(string(contents), -1)

	for _, match := range matches {
		varName := match[1]
		mu.Lock()
		varIndex[varName] = true
		mu.Unlock()
	}
}

func scanFiles(fileChan chan string, wg *sync.WaitGroup, extensions []string) {
	defer wg.Done()

	for filePath := range fileChan {
		scanFile(filePath, extensions)
	}
}

func checkEnv(directories []string, extensions []string, ignoreDirs []string, scanFiles func(chan string, *sync.WaitGroup, []string)) {
	fileChan := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go scanFiles(fileChan, &wg, extensions)
	}

	hasExtension := func(path string, extensions []string) bool {
		for _, ext := range extensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	}

	isIgnoredDir := func(path string, ignoreDirs []string) bool {
		for _, dir := range ignoreDirs {
			if strings.HasPrefix(path, dir) {
				return true
			}
		}
		return false
	}

	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Error: The specified directory '%s' does not exist.\n", dir)
			os.Exit(1)
		}

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && hasExtension(path, extensions) && !isIgnoredDir(path, ignoreDirs) {
				fileChan <- path
			}

			return nil
		})

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	close(fileChan)
	wg.Wait()

	for varName := range varIndex {
		if os.Getenv(varName) == "" {
			fmt.Printf("Missing variable: %s\n", varName)
		}
	}
}

func isValidFile(path string, extensions []string) bool {
	ext := filepath.Ext(path)
	for _, e := range extensions {
		if e == ext {
			return true
		}
	}
	return false
}

func main() {
	dirsFlag := flag.String("dirs", "", "Comma-separated list of directories to scan")
	extsFlag := flag.String("exts", "", "Comma-separated list of file extensions to scan")
	ignoreFlag := flag.String("ignore", "node_modules,vendor", "Comma-separated list of directories to ignore")
	flag.Parse()

	if *dirsFlag == "" || *extsFlag == "" {
		fmt.Println("Error: You must specify the directories and file extensions to scan using the -dirs and -exts flags.")
		os.Exit(1)
	}

	directories := strings.Split(*dirsFlag, ",")
	for i, dir := range directories {
		directories[i] = strings.TrimSpace(dir)
	}

	extensions := strings.Split(*extsFlag, ",")
	for i, ext := range extensions {
		extensions[i] = strings.TrimSpace(ext)
	}

	ignoreDirs := strings.Split(*ignoreFlag, ",")
	for i, dir := range ignoreDirs {
		ignoreDirs[i] = strings.TrimSpace(dir)
	}

	checkEnv(directories, extensions, ignoreDirs, scanFiles)
}
