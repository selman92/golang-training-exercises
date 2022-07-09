package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	listFilesFlag := flag.Bool("a", false, "List all flags: -a")
	flag.Parse()

	files := listFiles("testdata", listFilesFlag)

	for _, file := range files {
		fmt.Println(file)
	}
}

func listFiles(dir string, listAllFiles *bool) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileName := f.Name()

		if *listAllFiles == false && strings.HasPrefix(fileName, ".") {
			continue
		}

		dirs = append(dirs, fileName)
	}

	return dirs
}
