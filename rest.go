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