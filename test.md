


```
[jhuang@localhost rlf]$ go test -bench 10K

2020/12/20 18:49:32 running b.N = 1.
2020/12/20 18:49:35 4502638160 bytes read.
BenchmarkReadAsBytesWith10KBuf-4   	       1	2876863911 ns/op	4733158560 B/op	16806240 allocs/op
2020/12/20 18:49:35 running b.N = 1.
2020/12/20 18:49:37 4485832118 bytes read.
BenchmarkScanAsTextWith10KBuf-4    	       1	2479432486 ns/op	4715024712 B/op	16798604 allocs/op
2020/12/20 18:49:37 running b.N = 1.
2020/12/20 18:49:39 4485832118 bytes read.
BenchmarkScanAsBytesWith10KBuf-4   	       1	1160335612 ns/op	   72752 B/op	      17 allocs/op
PASS
ok  	rlf	6.520s

2020/12/20 18:49:50 running b.N = 1.
2020/12/20 18:49:53 4502638160 bytes read.
BenchmarkReadAsBytesWith10KBuf-4   	       1	2882788799 ns/op	4733158712 B/op	16806234 allocs/op
2020/12/20 18:49:53 running b.N = 1.
2020/12/20 18:49:56 4485832118 bytes read.
BenchmarkScanAsTextWith10KBuf-4    	       1	2505178502 ns/op	4715025992 B/op	16798621 allocs/op
2020/12/20 18:49:56 running b.N = 1.
2020/12/20 18:49:57 4485832118 bytes read.
BenchmarkScanAsBytesWith10KBuf-4   	       1	1163712152 ns/op	   72976 B/op	      19 allocs/op
PASS
ok  	rlf	6.555s
```