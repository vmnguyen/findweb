package main

import (
	"flag"
	"fmt"
)

func main() {
	f1 := flag.String("m", "domain", "Mode to scan: IP or Domain")
	f2 := flag.String("f", "hosts.txt", "Input file to scan")
	flag.Parse()
	scanMode := *f1
	scanFile := *f2
	fmt.Println("scanMode:", scanMode)
	fmt.Println("scanFile:", scanFile)
}


func probe(target string, port int) bool{
	// Probe probe here 
	return false
}