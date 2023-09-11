package crawl

const (
	DFS = "DFS"
	BFS = "BFS"
	Default = "BFS"
)

type Algorithm string

type Crawler interface {
	crawl(url string)
}

func NewCrawler(algorithm Algorithm) *Crawler {
	switch algorithm {
	case DFS:
		// return NewDFSCrawler()
	default:
		// return &BFSCrawler{}
	}
}