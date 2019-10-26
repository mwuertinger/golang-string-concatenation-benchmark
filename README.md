# Go string concatenation benchmark
Micro benchmark showing the performance of Sprint() vs. Sprintf() style string
concatenation.

## Build and run
```
go test -bench .
```

## Results
The results show that `Sprintf()` is actually faster:
```
BenchmarkSprintf-4      16268862                73.7 ns/op
BenchmarkSprint-4       12674227                94.7 ns/op
```
