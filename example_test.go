package synctools_test

import (
	"fmt"
	"github.com/gwik/synctools"
	"net/http"
	"time"
)

func ExamplePool() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://github.com/",
		"http://bitbucket.org/",
		"http://http://sigkill.tumblr.com/",
	}

	pool := synctools.NewPool(3)
	for _, url := range urls {
		// local copy of url. see http://golang.org/doc/effective_go.html#channels
		url := url
		pool.Spawn(func() {
			// Fetch the URL.
			fmt.Println("fetch", url)
			http.Get(url)
			time.Sleep(1 * time.Second)
		})
	}
	// Wait for all HTTP fetches to complete.
	pool.Wait()
}
