package crawl

type DFSCrawler struct {
	seeds[] string
	queue []string
	visited map[string]bool
}

func NewDFSCrawler(seeds []string) *DFSCrawler {
	return &DFSCrawler{
		seeds: seeds,
		queue: make([]string, 0),
		visited: make(map[string]bool),
	}
}

func (c *DFSCrawler) Crawl(url string) {
	c.queue = append(c.queue, url)

	for len(c.queue) > 0 {
		url := c.queue[0]
		c.queue = c.queue[1:]
		// do something
		
	}
}