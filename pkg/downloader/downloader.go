package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Downloader struct {}

func (d *Downloader) DownloadAndGetLinks(url string) [][]string {
	fmt.Println("Downloading", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading", url)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body", url)
	}

	html := string(content)
	var re = regexp.MustCompile(`(<img[^>]+src)="([^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)

	return matches
}

