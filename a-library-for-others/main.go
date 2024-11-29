package main

import (
	"bufio"
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
	line     string
	fields   []string
	lineRead bool
}

func (p csvParser) ReadLine(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	p.line = strings.TrimSuffix(line, "\n")
	p.fields = strings.Split(p.line, ",")
	p.lineRead = true
	return p.line, nil
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

	fmt.Println(parser.ReadLine(reader))
}
