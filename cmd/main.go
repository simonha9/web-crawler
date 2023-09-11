package cmd

import (

)

type Crawler struct {
	seeds []string
	algorithm string // currently only DFS or BFS
}

func (c *Crawler) Run() {
	for _, seed := range c.seeds {
		c.crawl(seed)
	}
}


func main() {
	c := Crawler{
		seeds: []string{},
	}
	c.Run()
}