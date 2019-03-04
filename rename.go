package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/schollz/progressbar"
)

func main() {
	globPattern := flag.String("glob", "", "glob of files to rename")
	outputDir := flag.String("output", "", "output directory")
	remove := flag.Bool("remove", false, "remove old files")
	outName := flag.String("name", "", "new name for files")
	max := flag.String("max", "1000", "max number for counting")
	flag.Parse()

	files, err := filepath.Glob(*globPattern)
	failOnErr(err)

	if len(files) == 0 {
		fmt.Println("no files globbed")
		os.Exit(0)
	}

	if *outputDir == "" {
		dir, _ := filepath.Split(files[0])
		fmt.Printf("Saving to default directory (input directory: %s)\n", dir)
		outputDir = &dir
	}

	length := strconv.Itoa(len(*max))

	var format string
	if *outName == "" {
		format = "%0" + length + "d"
	} else {
		format = *outName + "%0" + length + "d"
	}

	bar := progressbar.New(len(files))

	for idx, file := range files {
		_, name := filepath.Split(file)
		extension := filepath.Ext(name)

		newName := fmt.Sprintf(format, idx+1) + extension
		copy(file, *outputDir+newName)

		if *remove {
			err := os.Remove(file)
			if err != nil {
				fmt.Println(file + ": " + err.Error())
			}
		}
		bar.Add(1)
	}
	fmt.Println()
}

func failOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
