package main

import (
	"io"
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"bufio"
)

// CountLines - count the lines in from Reader
func CountLines(r io.Reader) (int, error) {

	return 0, nil
}

// CountFile write to stdout number of lines in the file

type Output []struct {
	FileInDir string
	LinesCount int
}




func CountFile(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var Count Output = make(Output,len(files))

		for i, f:=range files {
			Count[i].FileInDir = f.Name()
			Count[i].LinesCount = CountDir(path+"/"+f.Name())
	}
	fmt.Println(Count)
}

// CountDir - write to stdout numbers of lines for each files in directory
func CountDir(dir string) int {
	CountLinesInFile:=0
	files, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(files)
	for scanner.Scan() {
		CountLinesInFile++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return CountLinesInFile
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountFile(arg)

	}
}
