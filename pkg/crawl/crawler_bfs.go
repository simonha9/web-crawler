package crawl

import (
	"fmt"
	"strings"

	"github.com/simonha9/web-crawler/pkg/downloader"
)

type BFSCrawler struct {
	seeds   []string
	queue   []string
	visited map[string]bool
	downloader downloader.Downloader
}

func NewBFSCrawler(seeds []string) BFSCrawler {
	return BFSCrawler{
		seeds:   seeds,
		queue:   make([]string, 0),
		visited: make(map[string]bool),
	}
}

func (c BFSCrawler) Crawl(url string) {
	c.queue = append(c.queue, url)

	for len(c.queue) > 0 {
		url := c.queue[0]
		c.queue = c.queue[1:]

		if !c.IsVisited(url) {
			c.Visit(url)
		}

		links := c.downloader.DownloadAndGetLinks(url)
		for _, link := range links {
			for _, l := range link {
				if strings.HasPrefix(l, "/") {
					l = url + l
				}
				if !c.IsVisited(l) {
					c.queue = append(c.queue, l)
				}
			}
		}
	}
}

func (c BFSCrawler) Visit(url string) {
	c.visited[url] = true
	fmt.Println("Visited", url)
}

func (c BFSCrawler) IsVisited(url string) bool {
	return c.visited[url]
}