package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var numWorkers = runtime.NumCPU()

type fileInfo struct {
	fileName string
	Size     int64
}

type fileInfoPath struct {
	fileInfo
	filePath string
}

type duplicates struct {
	filePaths map[fileInfo][]string
}

func (fd *duplicates) save(fi fileInfo, fp string) {
	fd.filePaths[fi] = append(fd.filePaths[fi], fp)
}

func (fd *duplicates) getDup() <-chan []string {
	r := make(chan []string)
	go func() {
		for _, paths := range fd.filePaths {
			if len(paths) > 1 {
				r <- paths
			}
		}
		close(r)
	}()
	return r
}

func scanDir(wg *sync.WaitGroup, dirs <-chan string, files chan<- fileInfoPath) {
	for dir := range dirs {
		lsDir, err := os.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range lsDir {
			if !i.IsDir() {
				info, err := i.Info()
				if err != nil {
					log.Fatal(err)
				}
				files <- fileInfoPath{
					fileInfo{fileName: i.Name(), Size: info.Size()},
					filepath.Join(dir, i.Name()),
				}
			}

		}
	}
	wg.Done()
}

func saveFileInfo(d *duplicates, files <-chan fileInfoPath, done chan<- struct{}) {
	for f := range files {
		d.save(f.fileInfo, f.filePath)
	}
	done <- struct{}{}
}

func fScanDir(path string) <-chan []string {
	wg := &sync.WaitGroup{}
	dirsChan := make(chan string, 2*numWorkers)
	filesChan := make(chan fileInfoPath, 5*numWorkers)
	done := make(chan struct{})
	d := &duplicates{filePaths: make(map[fileInfo][]string)}

	go saveFileInfo(d, filesChan, done)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go scanDir(wg, dirsChan, filesChan)
	}

	dirQueue := make([]string, 1)
	dirQueue[0] = path
	dirsChan <- path

	for len(dirQueue) > 0 {
		dirPath := dirQueue[0]
		dirQueue = dirQueue[1:]

		lsDir, err := os.ReadDir(dirPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range lsDir {
			if e.IsDir() {
				fi, err := e.Info()
				if err != nil {
					log.Fatal(err)
				}
				newDirPath := filepath.Join(dirPath, fi.Name())

				dirQueue = append(dirQueue, newDirPath)
				dirsChan <- newDirPath
			}
		}
	}
	close(dirsChan)
	wg.Wait()
	close(filesChan)
	<-done
	return d.getDup()
}
