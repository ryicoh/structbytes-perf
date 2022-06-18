# structbytes-perf

I compared the performance of `unsafe` or `encoding/binary` in converting struct to bytes.

# Result

To use `unsafe` is faster than `encoding/binary`.

```
ryicoh@ryicohs-MacBook-Air structbytes-perf % go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/ryicoh/structbytes-perf
BenchmarkBinary-8                5215704               197.0 ns/op
BenchmarkUnsafeMarshal-8        1000000000               0.3931 ns/op
```
