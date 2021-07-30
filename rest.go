package gownload

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Download a file with Options.
func Download(url string, option Options) error {
	// Send a GET request.
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	body := resp.Body

	defer body.Close()

	// Read data from body.
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	// Create folder before Write.
	path := strings.Join(option.Folder, "/")
	os.MkdirAll(path, os.ModePerm)

	// Write data to the file.
	err = os.WriteFile(fmt.Sprintf("%s/%s", path, option.Name), data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Download Multiple file at same time.
// File names will be like photo_1.png, photo_2.png ...
func DownloadMultiple(urls []string, option Options) []error {
	channel := make(chan error)
	var results []error

	for index, url := range urls {
		go goroutineDownload(url, renameFile(option, index+1), channel)
	}

	for range urls {
		results = append(results, <-channel)
	}

	return results
}
