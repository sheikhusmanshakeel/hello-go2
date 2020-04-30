package categoriser

import (
	"fmt"
	"github.com/sheikhusmanshakeel/hello-go/pkg/crawler"
	"github.com/sheikhusmanshakeel/hello-go/pkg/indexer"
)
func Categorise(url string)  {
	fmt.Println(fmt.Sprintf("Categoriser called with URL: %s", url))
	crawler.Crawl()
	indexer.IndexURL(67)
}
