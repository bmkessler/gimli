# gimli

Gimli, a 384-bit permutation designed to achieve high
security with high performance across a broad range of platforms.

This is a golang port of the public domain reference implementation in C from:

https://gimli.cr.yp.to

## Benchmarks
```
Intel(R) Core(TM) i3-4010U CPU @ 1.70GHz
go version go1.9.2 linux/amd64

name     time/op
Gimli-4  488ns ± 1%
```
