


```bash
[jhuang@localhost rlf]$ go test -bench 100K -filetype utf16
goos: linux
goarch: amd64
pkg: rlf
BenchmarkReadAsBytesWith100KBuf-4   	       4	 275648977 ns/op	609696406 B/op	 1000027 allocs/op
BenchmarkScanAsTextWith100KBuf-4    	       4	 283194241 ns/op	609757894 B/op	 1000027 allocs/op
BenchmarkScanAsBytesWith100KBuf-4   	       9	 119894828 ns/op	  168392 B/op	      14 allocs/op
PASS
ok  	rlf	6.673s
[jhuang@localhost rlf]$ go test -bench 100K -filetype utf8
goos: linux
goarch: amd64
pkg: rlf
BenchmarkReadAsBytesWith100KBuf-4   	       5	 204294358 ns/op	419719115 B/op	 1000022 allocs/op
BenchmarkScanAsTextWith100KBuf-4    	       5	 204893402 ns/op	418735324 B/op	  999985 allocs/op
BenchmarkScanAsBytesWith100KBuf-4   	      12	  92065844 ns/op	  168328 B/op	      13 allocs/op
PASS
ok  	rlf	6.279s
[jhuang@localhost rlf]$ go test -bench 100K -filetype ascii
goos: linux
goarch: amd64
pkg: rlf
BenchmarkReadAsBytesWith100KBuf-4   	       6	 190196454 ns/op	364996430 B/op	 1000020 allocs/op
BenchmarkScanAsTextWith100KBuf-4    	       6	 192065210 ns/op	364010986 B/op	  999757 allocs/op
BenchmarkScanAsBytesWith100KBuf-4   	      12	  88537592 ns/op	  168403 B/op	      14 allocs/op
PASS
ok  	rlf	5.416s
[jhuang@localhost rlf]$ ll data/
total 11571452
-rw-rw-r--. 1 jhuang jhuang 4502638160 Dec 20 18:57 old-newspaper.ascii
-rw-rw-r--. 1 jhuang jhuang  346236861 Dec 20 19:28 old-newspaper-small.ascii
-rw-rw-r--. 1 jhuang jhuang  577358182 Dec 20 20:32 old-newspaper-small.utf16
-rw-rw-r--. 1 jhuang jhuang  398228796 Dec 20 20:12 old-newspaper-small.utf8
-rw-rw-r--. 1 jhuang jhuang 6024697599 May 12  2020 old-newspaper.tsv

```
