package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

func findFile(fileToSearch string, path string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(path)
	re := regexp.MustCompile("(?i)" + fileToSearch)
	if err != nil {
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if !entry.IsDir() {
			if re.MatchString(entry.Name()) {
				ch <- fullPath
			}
		} else {
			wg.Add(1)
			go findFile(fileToSearch, fullPath, ch, wg)
		}
	}
}

func main() {
	fileToSearch := flag.String("file", "", "Archivo a buscar")
	deepSearch := flag.Bool("deep", false, "Busca desde el directorio base")
	start := time.Now()

	flag.Parse()
	if *fileToSearch == "" {
		flag.Usage()
		os.Exit(1)
	}

	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	baseDir, _ := os.Getwd()

	if *deepSearch {
		baseDir = "/"
	}

	go findFile(*fileToSearch, baseDir, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()

	pathsFound := make([]string, 0)
	for path := range ch {
		pathsFound = append(pathsFound, path)
	}
	finish := time.Now()
	fmt.Println("Archivos encontrados: \n", len(pathsFound))
	for _, path := range pathsFound {
		fmt.Println(path)
	}
	fmt.Println("\nTiempo de ejecución: ", finish.Sub(start))
}
