package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

func findFile(fileToSearch string, path string, excludedRoutes []string, ch chan string, wg *sync.WaitGroup) {
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
			for _, excludedRoute := range excludedRoutes {
				if strings.EqualFold(entry.Name(), excludedRoute) {
					continue
				} else {
					wg.Add(1)
					go findFile(fileToSearch, fullPath, excludedRoutes, ch, wg)
				}
			}
		}
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("\nTiempo de ejecución: ", time.Since(start))
	}()

	fileToSearch := flag.String("file", "", "Archivo a buscar")
	deepSearch := flag.Bool("deep", false, "Busca desde el directorio base del sistema")
	excludeRouteFlag := flag.String("exclude", "", "Directorio a excluir, separados por comas sin espacios")
	helpFlag := flag.Bool("help", false, "Muestra la ayuda")

	flag.Parse()
	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if *fileToSearch == "" {
		fmt.Println("Debe especificar un archivo a buscar")
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println("Buscando archivos...")

	var excludedRoutes []string
	excludedRoutes = strings.Split(*excludeRouteFlag, ",")
	fmt.Println(excludedRoutes)
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	baseDir, _ := os.Getwd()

	if *deepSearch {
		baseDir = "/"
	}

	go findFile(*fileToSearch, baseDir, excludedRoutes, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()

	pathsFound := make([]string, 0)
	for path := range ch {
		pathsFound = append(pathsFound, path)
	}
	fmt.Println("\n", len(pathsFound), "archivos encontrados")
	for _, path := range pathsFound {
		dir := filepath.Dir(path)
		link := fmt.Sprintf("\033]8;;file://%s\033\\%s\033]8;;\033\\", dir, path)
		fmt.Println(link)
	}
}
