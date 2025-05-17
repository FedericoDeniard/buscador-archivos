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
	fmt.Println("Buscando archivos...")
	start := time.Now()
	fileToSearch := flag.String("file", "", "Archivo a buscar")
	deepSearch := flag.Bool("deep", false, "Busca desde el directorio base del sistema")

	flag.Parse()
	if *fileToSearch == "" {
		fmt.Println("Debe especificar un archivo a buscar")
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
	fmt.Println("\n", len(pathsFound), "archivos encontrados")
	for _, path := range pathsFound {
		dir := filepath.Dir(path)
		link := fmt.Sprintf("\033]8;;file://%s\033\\%s\033]8;;\033\\", dir, path)
		fmt.Println(link)
	}
	fmt.Println("\nTiempo de ejecución: ", finish.Sub(start))
}
