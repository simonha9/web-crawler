package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Downloader struct {
	content map[string]string
}

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

	hash := d.hash(html)
	if _, ok := d.content[hash]; ok {
		return nil
	}

	var re = regexp.MustCompile(`(<img[^>]+src)="([^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)

	return matches
}

func (d *Downloader) hash(html string) string {
	hasher := md5.New()
	hasher.Write([]byte(html))
	return hex.EncodeToString(hasher.Sum(nil))
}