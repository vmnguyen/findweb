package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/valyala/fasthttp"
)

var (
	client = &fasthttp.Client{
		ReadTimeout:              time.Second,
		NoDefaultUserAgentHeader: true,
		MaxConnsPerHost:          10000,
	}
	ports = []string{"80", "443", "8000", "8080"}
)

func doRequest(i interface{}) {
	url := i.(string)
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	resp.SkipBody = true
	//check if url has http/https
	req.SetRequestURI(url)
	client.Do(req, resp)
	bodyByte := resp.StatusCode()
	fmt.Println("[", bodyByte, "] ", url)
}

func scanPort(filePath string, mode bool) {
	defer ants.Release()

	runTimes := 100

	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(runTimes, func(i interface{}) {

		doRequest(i)
		wg.Done()
	})
	defer p.Release()

	if mode {
		// Scan by domain
		fmt.Println("[~] Using list of domain")
		file, _ := os.Open(filePath)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			url := scanner.Text()
			if !strings.Contains(url, "http") {
				url = "https://" + url
			}
			//fmt.Println("Full url: ", url)

			wg.Add(1)
			_ = p.Invoke(string(url))
			//doRequest(url)

		}
		wg.Wait()
		fmt.Printf("running goroutines: %d\n", p.Running())
	} else {
		// Scan by IP address
		fmt.Println("[~] Using list of IP address")
	}
}

func myFunction(i interface{}) {
	url := i.(string)
	fmt.Println("Running process", url)
}

func main() {

	var mode string
	flag.StringVar(&mode, "m", "domain", "Input type: Domain or IP?")

	var filePath string
	flag.StringVar(&filePath, "f", "hosts.txt", "Filepath")

	var concurrency int
	flag.IntVar(&concurrency, "c", 50, "Concurrentcy")

	flag.Parse()

	if mode == "domain" {
		scanPort(filePath, true)
	} else {
		scanPort(filePath, false)
	}

}
func probe(target string, port int) bool {
	// Probe probe here
	return false
}
