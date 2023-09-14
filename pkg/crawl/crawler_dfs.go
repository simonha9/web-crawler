package crawl

import (
	"fmt"
	"strings"

	"github.com/simonha9/web-crawler/pkg/downloader"
)

const MAX_DEPTH = 50

type DFSCrawler struct {
	seeds   []string
	visited map[string]bool
	downloader downloader.Downloader
}

func NewDFSCrawler(seeds []string) *DFSCrawler {
	return &DFSCrawler{
		seeds:   seeds,
		visited: make(map[string]bool),
	}
}

func (c *DFSCrawler) Crawl(url string) {
	c.crawl(url, 0)
}

func (c *DFSCrawler) crawl(url string, depth int) {
	if c.IsVisited(url) {
		return
	}

	if depth > MAX_DEPTH {
		return
	}
	
	c.Visit(url)

	links := c.downloader.DownloadAndGetLinks(url)
	for _, link := range links {
		for _, l := range link {
			if strings.HasPrefix(l, "/") {
				l = url + l
			}
			c.crawl(l, depth + 1)
		}
	}
} 

func (c DFSCrawler) Visit(url string) {
	c.visited[url] = true
	fmt.Println("Visited", url)
}

func (c DFSCrawler) IsVisited(url string) bool {
	return c.visited[url]
}