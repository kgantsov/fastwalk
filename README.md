# Fastwalk

This repository a bit rewritten version of fastwalk package from github.com/golang/tools/tree/master/internal/fastwalk

See https://github.com/golang/tools/tree/master/internal/fastwalk.


### Example:

    package main

    import (
        "fmt"
        "os"
        "github.com/kgantsov/fastwalk"
    )

    func walkFunction(path string, fileType os.FileMode) error {
        if fileType == os.ModeDir {
            fmt.Printf("Found directory: %s \n", path)
        } else {
            fmt.Printf("Found file: %s \n", path)
        }
        return nil
    }

    func main() {
        err := fastwalk.Walk(dir, walkFunction)

        if err != nil {
            fmt.Printf("Walking error: %s", err)
        }
    }


### Benchmarks

    go test -run XXX -bench . -benchmem

    goos: darwin
    goarch: amd64
    pkg: github.com/kgantsov/fastwalk
    BenchmarkFastWalk-4   	      30	  56641299 ns/op	 2297299 B/op	   27086 allocs/op
    BenchmarkWalk-4       	      10	 130804203 ns/op	17779115 B/op	   54873 allocs/op
    PASS
    ok  	github.com/kgantsov/fastwalk	5.472s
