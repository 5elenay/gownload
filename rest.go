package gownload

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Download a file with Options.
func Download(url string, option Options) Result {
	// Send a GET request.
	resp, err := http.Get(url)

	if err != nil {
		return Result{
			Error: err,
			Path: "",
		}
	}

	body := resp.Body

	defer body.Close()

	// Read data from body.
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return Result{
			Error: err,
			Path: "",
		}
	}

	// Create folder before Write.
	path := strings.Join(option.Folder, "/")
	os.MkdirAll(path, os.ModePerm)

	// Write data to the file.
	filePath := strings.Join([]string{path, option.Name}, "/")
	err = os.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return Result{
			Error: err,
			Path: "",
		}
	}

	return Result{
			Error: nil,
			Path: filePath,
		}
}

// Download Multiple file at same time.
// File names will be like photo_1.png, photo_2.png ...
func DownloadMultiple(urls []string, option Options) []Result {
	channel := make(chan Result)
	var results []Result

	for index, url := range urls {
		go goroutineDownload(url, renameFile(option, index+1), channel)
	}

	for range urls {
		results = append(results, <-channel)
	}

	return results
}
