# Files trasversers logic

**current.go:** Represents manual trasverse method by using `os.ReadDir` and recursion for sub-directories. The recursion uses goroutine.

**walker.go:** Implements the usage of `filepath.WalkDir` which is an inbuit trasverse mechanism.

## BenchMarkTest

This repo root is used as the target entry point. See `manual_vs_filepath_walk_test.go`

```go
  goos: linux
  goarch: amd64
  cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
  BenchmarkFilepathWalDir-4          25618             47425 ns/op
  BenchmarkManualRead-4              40051             29239 ns/op
  PASS
  ok      command-line-arguments  3.171s
```

The manual trasverse with recursion and goroutine is more faster than the filepath.WalkDir method.

Note: Please conduct the test on your machine as well to kknow the best method to use.
