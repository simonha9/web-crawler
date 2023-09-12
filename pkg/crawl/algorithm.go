package crawl

const (
	DFS Algorithm = "DFS"
	BFS Algorithm = "BFS"
	Default Algorithm = "BFS"
)

type Algorithm string

type Crawler interface {
	Crawl(url string)
}

func NewCrawler(algorithm Algorithm, seeds []string) Crawler {
	switch algorithm {
	case DFS:
		// return NewDFSCrawler()
	default:
		return NewBFSCrawler(seeds)
	}
	
	return nil
}