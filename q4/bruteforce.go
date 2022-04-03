
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func process(url string, ch chan time.Duration) {
	t1 := time.Now()

	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = body

	_ = nbytes
	_ = err

	t2 := time.Now()

	diff := t2.Sub(t1)

	ch <- diff

}

func main() {
	tstart := time.Now()

	no := 3
	ch := make(chan time.Duration, 3)
	addresses := [...]string{"https://www.python.org/", "https://www.ruby-lang.org/en/", "https://golang.org/"}

	for i := 0; i < no; i++ {
		process(addresses[i], ch)
	}

	fmt.Println("read value", <-ch)
	fmt.Println("read value", <-ch)
	fmt.Println("read value", <-ch)

	tend := time.Now()

	totaldiff := tend.Sub(tstart)

	fmt.Println("Total Time taken: ", totaldiff)

}
