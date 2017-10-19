# gohumantime
A human readable time translator in golang

[![GoDoc](https://godoc.org/github.com/urjitbhatia/gohumantime?status.svg)](https://godoc.org/github.com/urjitbhatia/gohumantime)

# Motivation
Use a simple human readable string representation of time
Inspired by: https://github.com/rschmukler/human-interval

# Usage:

```go
import "github.com/urjitbhatia/gohumantime"
//...
millis, err := gohumantime.ToMilliseconds("1 day and 14 hours")

millis, err := gohumantime.ToMilliseconds("9981234314")
// If string is numeric, unit is assumed to be Milliseconds and a simple atoi conversion is returned.
//...
```

# Benchmarks
```go
  //3000 samples of ToMilliseconds("4 seconds, 2 minutes 1 hour and 3 days and 10weeks, 1 month and 1 year"):

  Ran 3000 samples:
  runtime:
    Fastest Time: 0.000s
    Slowest Time: 0.001s
    Average Time: 0.000s ± 0.000s
```
