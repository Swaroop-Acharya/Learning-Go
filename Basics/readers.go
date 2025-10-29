package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func stringReader() {
	s := strings.NewReader("This is a string yo")
	readerToStdOut(s,2)
}

func fileReader() {
	f, err := os.Open("../Resources/Notes.txt")
	if err != nil {
		fmt.Println("Error:",err)	
	}
	defer f.Close()
	readerToStdOut(f,10)
}

func readerToStdOut(r io.Reader, buffSize int){
	buff := make([]byte, buffSize)
	for {
		n, err := r.Read(buff)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buff[:n]))
		}
	}
}

func main() {
	stringReader()
	fileReader()
}
