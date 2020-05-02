package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

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

func doRequest(url string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(url)
	client.Do(req, resp)
	bodyByte := resp.Body()
	fmt.Println(string(bodyByte))
}

func scanPort(filePath string, mode bool) {
	if mode {
		// Scan by domain
		fmt.Println("[~] Using list of domain")
		file, _ := os.Open(filePath)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			doRequest(scanner.Text())
		}

	} else {
		// Scan by IP address
		fmt.Println("[~] Using list of IP address")
	}
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
