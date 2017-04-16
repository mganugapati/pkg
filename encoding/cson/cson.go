package cson

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/scanner"

	"github.com/mkideal/pkg/encoding"
)

// cson represents Comma-Separated Object Notation
//
// types:
//	Integer
//	Float
//	String
//	Object
//	Array
//
// examples:
//
// 1
// 1,2
// 1,2.3
// "1",2
// true,false
// 1,{2,3},["a","b"],{"c"},[1,2,3]

type lineReader struct {
	reader  io.Reader
	eof     bool
	lineeof bool
}

func (r *lineReader) reset() {
	r.lineeof = false
}

func (r *lineReader) Read(p []byte) (int, error) {
	if r.lineeof {
		return 0, io.EOF
	}
	size := len(p)
	readedNum := 0
	// read bytes one by one util EOF or '\n' readed
	for i := 0; i < size; i++ {
		n, err := r.reader.Read(p[i : i+1])
		readedNum += n
		r.eof = r.eof || err == io.EOF
		if err != nil {
			return readedNum, err
		}
		if p[i] == '\n' {
			r.lineeof = true
			return readedNum, io.EOF
		}
	}
	return readedNum, nil
}

func newScanner() *scanner.Scanner {
	s := new(scanner.Scanner)
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings
	return s
}

func readLine(p *parser, s *scanner.Scanner, r *lineReader) (Node, error) {
	s = s.Init(r)
	if err := p.init(s); err != nil {
		return nil, err
	}
	return p.parseListNode(encoding.ObjectNode, '{', '}', false)
}

// Read reads one node
func Read(r io.Reader) (Node, error) {
	s := newScanner()
	p := new(parser)
	lr := &lineReader{reader: r}
	return readLine(p, s, lr)
}

// ReadAll reads all nodes
func ReadAll(r io.Reader) ([]Node, error) {
	s := newScanner()
	p := new(parser)
	lr := &lineReader{reader: r}
	var nodes []Node
	for !lr.eof {
		lr.reset()
		n, err := readLine(p, s, lr)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
		fmt.Printf("lr.eof: %v\n", lr.eof)
	}
	return nodes, nil
}

// ReadBytes read a json node from bytes
func ReadBytes(data []byte) ([]Node, error) {
	return ReadAll(bytes.NewBuffer(data))
}

// ReadFile read a json node from file
func ReadFile(filename string) ([]Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadAll(file)
}
