# SaCache

SaCache(a.k.a Sashi Cache) is a fast, concurrent in-memory cache service written in pure Go. It uses gRPC for communicating with clients and supports item expiration function. It is still under development. The goal of SaCache is reducing extra costs(GC) as much as possible and making it faster and faster.

## Usage

Please read [sacache_test.go](sacache_test.go) for lib usage purpose and read [server.go](server/server.go) & [client.go](client/client.go) for C-S usage purpose.

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/sashirin/sacache
cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
BenchmarkBigCacheSet-6      	     225	   5109367 ns/op	  12.83 MB/s	 4180315 B/op	      50 allocs/op
BenchmarkBigCacheGet-6      	     374	   3048027 ns/op	  21.50 MB/s	 1424270 B/op	  131099 allocs/op
BenchmarkBigCacheSetGet-6   	     146	   8126713 ns/op	  16.13 MB/s	 2829722 B/op	  131142 allocs/op
BenchmarkSaCacheSet-6       	     150	   7922742 ns/op	   8.27 MB/s	 5560263 B/op	  196680 allocs/op
BenchmarkSaCacheGet-6       	     631	   1864814 ns/op	  35.14 MB/s	  283983 B/op	   65864 allocs/op
BenchmarkSaCacheSetGet-6    	     181	   9285902 ns/op	  14.12 MB/s	 5812962 B/op	  262204 allocs/op
BenchmarkStdMapSet-6        	     152	   8064075 ns/op	   8.13 MB/s	  344674 B/op	   65551 allocs/op
BenchmarkStdMapGet-6        	     450	   2408782 ns/op	  27.21 MB/s	   28471 B/op	     150 allocs/op
BenchmarkStdMapSetGet-6     	      82	  37934010 ns/op	   3.46 MB/s	  415437 B/op	   65565 allocs/op
BenchmarkSyncMapSet-6       	      72	  16459389 ns/op	   3.98 MB/s	 3537239 B/op	  263997 allocs/op
BenchmarkSyncMapGet-6       	    1557	    743652 ns/op	  88.13 MB/s	    8169 B/op	     254 allocs/op
BenchmarkSyncMapSetGet-6    	     343	   3145552 ns/op	  41.67 MB/s	 3434993 B/op	  262532 allocs/op
```

## License

SaCache is released under the MIT license (see [LICENSE](LICENSE))
