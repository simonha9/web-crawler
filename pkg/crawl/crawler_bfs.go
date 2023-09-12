package crawl

import "fmt"

type BFSCrawler struct {
	seeds   []string
	queue   []string
	visited map[string]bool
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

		fmt.Println(url)
		// do something

	}
}