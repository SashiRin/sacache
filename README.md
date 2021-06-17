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
BenchmarkBigCacheSet-6      	     237	   5311201 ns/op	  12.34 MB/s	 3968662 B/op	      47 allocs/op
BenchmarkBigCacheGet-6      	     379	   3199278 ns/op	  20.48 MB/s	 1412395 B/op	  131099 allocs/op
BenchmarkBigCacheSetGet-6   	     147	   8616085 ns/op	  15.21 MB/s	 2814039 B/op	  131141 allocs/op
BenchmarkSaCacheSet-6       	     136	  10319922 ns/op	   6.35 MB/s	 6685742 B/op	  131294 allocs/op
BenchmarkSaCacheGet-6       	     628	   1912921 ns/op	  34.26 MB/s	  284503 B/op	   65773 allocs/op
BenchmarkSaCacheSetGet-6    	     144	  14731054 ns/op	   8.90 MB/s	 6998752 B/op	  196820 allocs/op
BenchmarkStdMapSet-6        	     150	   7829785 ns/op	   8.37 MB/s	  345839 B/op	   65551 allocs/op
BenchmarkStdMapGet-6        	     412	   2840184 ns/op	  23.07 MB/s	   31084 B/op	     164 allocs/op
BenchmarkStdMapSetGet-6     	      81	  43329190 ns/op	   3.03 MB/s	  417120 B/op	   65565 allocs/op
BenchmarkSyncMapSet-6       	      74	  16229905 ns/op	   4.04 MB/s	 3533916 B/op	  263947 allocs/op
BenchmarkSyncMapGet-6       	    1533	    772224 ns/op	  84.87 MB/s	    8302 B/op	     258 allocs/op
BenchmarkSyncMapSetGet-6    	     334	   3408553 ns/op	  38.45 MB/s	 3435798 B/op	  262543 allocs/op
```

## License

SaCache is released under the MIT license (see [LICENSE](LICENSE))
