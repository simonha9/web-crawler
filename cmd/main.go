package cmd

import (
	"github.com/simonha9/web-crawler/pkg/frontier"
	"flag"
)

// write code below to take the algorithm as a flag with --algo or -a
var (
	algorithm = flag.String("algo", "BFS", "BFS or DFS")
)


type Crawler struct {
	seeds []string
	algorithm string // currently only DFS or BFS
}

func (c *Crawler) Run() {
	frontier := frontier.NewFrontier(c.seeds, c.algorithm)
	frontier.Run()
}


func main() {
	flag.Parse()

	c := Crawler{
		seeds: []string{},
		algorithm: *algorithm,
	}
	c.Run()
}