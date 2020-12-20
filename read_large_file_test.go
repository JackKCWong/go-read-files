package rlf

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"
)

var filetype string

func init() {
	flag.StringVar(&filetype, "filetype", "utf8", "specify the file type to load. ascii|utf8|utf16")
}

type testfn func(*bufio.Reader, int) (int, int)

func bench(b *testing.B, filepath string, bufsize int, fn testfn) {
	log.SetOutput(ioutil.Discard)
	log.Printf("running b.N = %d.", b.N)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		of, err := os.Open(filepath)

		if err != nil {
			b.Fatalf("failed to read file: %s", filepath)
		}

		charCount, lineCount := fn(bufio.NewReaderSize(of, bufsize*1024), 0)

		of.Close()
		log.Printf("%d bytes / %d lines read from %s.", charCount, lineCount, filepath)
		// memstat()
	}
}

func memstat() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("\n")
	fmt.Printf("Alloc: %d MB, TotalAlloc: %d MB, Sys: %d MB\n",
		ms.Alloc/1024/1024, ms.TotalAlloc/1024/1024, ms.Sys/1024/1024)
	fmt.Printf("Mallocs: %d, Frees: %d\n",
		ms.Mallocs, ms.Frees)
	fmt.Printf("HeapAlloc: %d MB, HeapSys: %d MB, HeapIdle: %d MB\n",
		ms.HeapAlloc/1024/1024, ms.HeapSys/1024/1024, ms.HeapIdle/1024/1024)
	fmt.Printf("HeapObjects: %d\n", ms.HeapObjects)
	fmt.Printf("\n")
}

func BenchmarkReadAsBytesWith1KBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 1, ReadAsBytes)
}

func BenchmarkReadAsBytesWith10KBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 10, ReadAsBytes)
}

func BenchmarkReadAsBytesWith100KBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 100, ReadAsBytes)
}

func BenchmarkReadAsBytesWith1MBBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 1024, ReadAsBytes)
}

func BenchmarkReadAsBytesWith10MBBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 10*1024, ReadAsBytes)
}

func BenchmarkScanAsTextWith100KBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 100, ScanAsText)
}

func BenchmarkScanAsBytesWith100KBuf(b *testing.B) {
	bench(b, "data/old-newspaper-small."+filetype, 100, ScanAsBytes)
}
