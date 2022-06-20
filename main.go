package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	pathFlag = flag.String("path", ".", "directory for search duplicates files")
	delFlag  = flag.Bool("d", false, "delete duplicates files")
)

func main() {
	flag.Parse()
	dupFiles := fScanDir(*pathFlag)

	for dupFile := range dupFiles {
		fmt.Printf("[%s] (%d):\n", filepath.Base(dupFile[0]), len(dupFile))
		for _, df := range dupFile {
			fmt.Println("\t", df)
		}
		if *delFlag {
			fmt.Print("Delete duplicates? (y/N) ")
			var ans string

			fmt.Scanln(&ans)

			if strings.ToLower(ans) == "y" {
				for _, f := range dupFile[1:] {
					err := os.Remove(f)
					if err != nil {
						fmt.Printf("Couldn't remove file: %v", err)
						os.Exit(1)
					}
				}
			}
		}
	}
}
