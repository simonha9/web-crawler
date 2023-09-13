package downloader

import "fmt"

type Downloader struct {}

func (d *Downloader) Download(url string) {
	fmt.Println("Downloading", url)
}
