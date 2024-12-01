package main

import (
	"fmt"
	"io"
	"strings"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

type csvParser struct {
	lastLine string
	fields   [][]string
	lineRead bool
	name     string
	age      int
}

func (p csvParser) ReadLine(r io.Reader) (string, error) {
	b := make([]byte, 1)
	line := ""

	for {
		_, err := r.Read(b)
		if err != nil {
			return "", err
		}

		c := string(b[0])
		if c == "\n" {
			break
		}

		line += c
	}

	return line, nil
}

func (mp csvParser) GetField(n int) (string, error) {
	return "", nil
}

func (mp csvParser) GetNumberOfFields() int {
	return 0
}

func main() {
	var parser CSVParser
	parser = csvParser{}

	reader := strings.NewReader("kek\nasd")
	reader2 := strings.NewReader("123\n456")

	fmt.Println(parser.ReadLine(reader))
	fmt.Println(parser.ReadLine(reader2))
}
