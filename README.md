# Fastwalk

This repository a bit rewritten version of fastwalk package from github.com/golang/tools/tree/master/internal/fastwalk

See https://github.com/golang/tools/tree/master/internal/fastwalk.


Example:

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
