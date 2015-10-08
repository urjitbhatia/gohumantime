# gohumantime
A human readable time translator in golang

[![GoDoc](https://godoc.org/github.com/urjitbhatia/gohumantime?status.svg)](https://godoc.org/github.com/urjitbhatia/gohumantime)

# Motivation
Use a simple human readable string representation of time
Inspired by: https://github.com/rschmukler/human-interval

# Benchmarks
```
  Slowest iteration is faster than 600 ns (Benchmark test fails if it takes more than 600 ns)
  2000 samples of ToMilliseconds("4 seconds, 2 minutes 1 hour and 3 days and 10weeks, 1 month and 1 year"):
      runtime:
        Fastest Time: 0.000s
        Slowest Time: 0.000s
        Average Time: 0.000s ± 0.000s
```
