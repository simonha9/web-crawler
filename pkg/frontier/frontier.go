package frontier

import (
	"github.com/simonha9/web-crawler/pkg/crawl"
)

type Frontier struct {
	seeds []string
	algorithm crawl.Algorithm // currently only DFS or BFS
	crawler crawl.Crawler
}


func NewFrontier(seeds []string, algorithm string) *Frontier {
	algo := crawl.Default
	switch algorithm {
	case "DFS":
		algo = crawl.DFS
	case "BFS":
		algo = crawl.BFS
	}

	return &Frontier{
		seeds: seeds,
		algorithm: algo,
		crawler: crawl.NewCrawler(algo, seeds),
	}
}

func (f *Frontier) Run() {
	for _, seed := range f.seeds {
		f.crawler.Crawl(seed)
	}
}

