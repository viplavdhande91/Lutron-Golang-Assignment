package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"sync"

)

var totalBytes int =0;

func process(url string, wg *sync.WaitGroup,ch chan time.Duration) {
	t1 := time.Now()

 	resp, err := http.Get(url)
 	body, err := ioutil.ReadAll(resp.Body)
 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    
	_ = nbytes
	
	 _ = err

	t2 := time.Now()

	totalBytes += len(body)

 	diff := t2.Sub(t1)

	wg.Done()
	ch <- diff

}

func main() {
	t1 := time.Now()

	no := 3
	ch := make(chan time.Duration, 3)
 	addresses := [...]string{"https://www.python.org/", "https://www.ruby-lang.org/en/", "https://golang.org/"}

	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(addresses[i], &wg,ch)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	fmt.Println("read value", <-ch)
    fmt.Println("read value", <-ch)
    fmt.Println("read value", <-ch)

	t2 := time.Now()

 	diff := t2.Sub(t1)

	fmt.Println("Total Time taken: ",diff)

	fmt.Println("Total Bytes: ",totalBytes)


}
