package rlf

import (
	"bufio"
	"io"
)

func ReadAsBytes(r *bufio.Reader, bufsize int) (int, int) {
	nr := bufio.NewReaderSize(r, bufsize)

	var err error
	var line []byte
	lineCount := 0
	sum := 0
	for err != io.EOF {
		line, err = nr.ReadBytes('\n')
		if err != nil {
			break
		}

		sum += len(line) - 1
		lineCount++
	}

	return sum, lineCount
}

func ScanAsText(r *bufio.Reader, bufsize int) (int, int) {
	nr := bufio.NewReaderSize(r, bufsize)
	scanner := bufio.NewScanner(nr)

	lineCount := 0
	sum := 0
	for scanner.Scan() {
		sum += len(scanner.Text())
		lineCount++
	}

	return sum, lineCount
}

func ScanAsBytes(r *bufio.Reader, bufsize int) (int, int) {

	sum := 0
	lineCount := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		sum += len(scanner.Bytes())
		lineCount++
	}

	return sum, lineCount
}
