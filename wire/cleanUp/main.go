package main

import (
	"fmt"
	"log"
	"os"
)

type FileReader struct {
	f *os.File
}

func NewFileReader(filePath string) (*FileReader, func(), error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	fr := &FileReader{
		f: f,
	}

	fn := func() {
		log.Println("clean up")
		fr.f.Close()
	}

	return fr, fn, nil
}

func main() {
	fr, cleanup, err := InitFileReader("./test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanup()

	buf := make([]byte, 1024)
	_, _ = fr.f.Read(buf)

	fmt.Println(string(buf))
}
