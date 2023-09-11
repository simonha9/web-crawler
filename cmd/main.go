package cmd

import (
	"github.com/simonha9/web-crawler/pkg/frontier"
	"github.com/simonha9/web-crawler/pkg/crawl"
)

type Crawler struct {
	seeds []string
	algorithm crawl.Algorithm // currently only DFS or BFS
}

func (c *Crawler) Run() {
	frontier := frontier.Frontier{c.seeds, c.algorithm}
	frontier.Run()
}


func main() {
	c := Crawler{
		seeds: []string{},
	}
	c.Run()
}