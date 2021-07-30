# Gownload

Go version of [pewn](https://github.com/5elenay/pewn). Allows you to Download file(s) easily.

## Installation

- Initialize your project with `go mod init <name>`.
- Get the package with `go get github.com/5elenay/gownload`.

## API Reference

Click [here](https://pkg.go.dev/github.com/5elenay/gownload) for API reference.

## Example

Download single file example

```go
package main

import (
    "fmt"

    "github.com/5elenay/gownload"
)

func main() {
    // Create an option
    option := gownload.Options{
        Name:   "photo.png",
        Folder: []string{"path", "to", "photos"},
    }

    // Download file.
    res := gownload.Download("https://picsum.photos/500/300", option)

    // An helper function for gownload.Result which allows you to handle result and error easily.
    res.Handle(res.Error, func(path string) {
        // Use file path
        fmt.Println(path)
    })
}
```

Download multiple-file example

```go
package main

import (
    "fmt"

    "github.com/5elenay/gownload"
)

func main() {
    // Create an option.
    option := gownload.Options{
        Name:   "photo.png",
        Folder: []string{"path", "to", "photos"},
    }

    // Make a string slice and add the url 10 times.
    var urls []string

    for i := 0; i < 10; i++ {
        urls = append(urls, "https://picsum.photos/500/300")
    }

    // Download all of the files.
    results := gownload.DownloadMultiple(urls, option)

    // Loop through files
    for _, res := range results {
        // An helper function for gownload.Result which allows you to handle result and error easily.
        res.Handle(res.Error, func(path string) {
            // Use file path
            fmt.Println(path)
        })
    }
}

```
