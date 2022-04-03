package main
//status q3:1 DONE , q3:2 NOT DONE
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Lang struct {
	Name          string
	URL           string
	bytes         uint
	time_duration time.Duration
}

func getNumberOfBytes(url string) (int, time.Duration) {
	t1 := time.Now()

	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	//fmt.Println(string(body))

	fmt.Println("Error: ", err)
	fmt.Println()
	_ = nbytes

	t2 := t1.Add(time.Second)

	diff := t2.Sub(t1)

	return len(body), diff
}

func crawl_output(lang Lang) {

	fmt.Println("CRAWLED OUTPUT FOR: ", lang.Name)
	fmt.Println("URL: ", lang.URL)
	fmt.Println("Bytes: ", lang.bytes)
	fmt.Println("Time Duration: ", lang.time_duration)
	fmt.Println()
}
func main() {

	addresses := [...]string{"https://www.python.org/", "https://www.ruby-lang.org/en/", "https://golang.org/"}

	addr1 := Lang{}
	addr2 := Lang{}
	addr3 := Lang{}

	//STRUCTURE VARIABLE 1
	addr1.Name = "Python"
	addr1.URL = addresses[0]
	nbytes, time := getNumberOfBytes(addr1.URL)
	addr1.bytes = uint(nbytes)
	addr1.time_duration = time

	//STRUCTURE VARIABLE 2
	addr2.Name = "Ruby"
	addr2.URL = addresses[1]
	nbytes, time = getNumberOfBytes(addr2.URL)
	addr2.bytes = uint(nbytes)
	addr2.time_duration = time

	//STRUCTURE VARIABLE 3
	addr3.Name = "Golang"
	addr3.URL = addresses[2]
	nbytes, time = getNumberOfBytes(addr3.URL)
	addr3.bytes = uint(nbytes)
	addr3.time_duration = time

	struct_var := [...]Lang{addr1, addr2, addr3}

	for i := 0; i < len(struct_var); i++ { //PRINT FORMATTED OUTPUT
		crawl_output((struct_var[i]))
	}

}
