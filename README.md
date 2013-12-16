`synctools` provides a Pool structure that limit
the concurrency of goroutines.

`Pool` wraps a `sync.WaitGroup` and uses a channels to limit the number of
goroutines to be run.

Example :

```go

package main

import (
  "fmt"
  "github.com/gwik/synctools"
  "net/http"
  "time"
)

func main() {
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
    // avoid for loop cache, see http://golang.org/doc/effective_go.html#channels
    url := url
    pool.Spawn(func() {
      fmt.Println("fetch", url)
      // Fetch the URL.
      http.Get(url)
      time.Sleep(3 * time.Second)
    })
  }
  // Wait for all HTTP fetches to complete.
  pool.Wait()
}

```