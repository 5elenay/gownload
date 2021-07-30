package gownload

import (
	"fmt"
	"strings"
)

func goroutineDownload(url string, option Options, channel chan Result) {
	channel <- Download(url, option)
}

func renameFile(option Options, number int) Options {
	splitted := strings.Split(option.Name, ".")
	var splitIndex int

	if len(splitted) == 1 {
		splitIndex = 0
	} else {
		splitIndex = len(splitted) - 2
	}

	fileName := splitted[splitIndex]

	splitted[splitIndex] = fmt.Sprintf("%s_%d", fileName, number)

	option.Name = strings.Join(splitted, ".")

	return option
}
