package frontier

import "fmt"

type Frontier struct {
	seeds []string
	algorithm string // currently only DFS or BFS
}

func (f *Frontier) Run() {
	for _, seed := range f.seeds {
		f.crawl(seed)
	}
}

func (f *Frontier) crawl(url string) {
	// do something
	switch f.algorithm {
	case "DFS":
		f.DFS()
	case "BFS":
		f.BFS()
	}
}

func (f *Frontier) DFS() {
	fmt.Println("DFS")
}

func (f *Frontier) BFS() {
	fmt.Println("BFS")
}


