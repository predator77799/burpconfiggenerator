package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Target struct {
		Scope struct {
			Include []Entry `json:"include"`
			Exclude []Entry `json:"exclude"`
		} `json:"scope"`
	} `json:"target"`
}

type Entry struct {
	Enabled  bool   `json:"enabled"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	File     string `json:"file,omitempty"`
}

func main() {
	// Parse command-line flags
	inputFile := flag.String("input", "scope.txt", "Path to the input file containing IP addresses")
	outputFile := flag.String("output", "burp_config.json", "Path to the output Burp Suite configuration file")
	flag.Parse()

	data, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	ips := strings.Split(strings.TrimSpace(string(data)), "\n")

	var config Config

	var wg sync.WaitGroup
	wg.Add(len(ips))

	for _, ip := range ips {
		go func(ip string) {
			defer wg.Done()

			config.Target.Scope.Include = append(config.Target.Scope.Include, Entry{
				Enabled:  true,
				Protocol: "any",
				Host:     ip,
			})
		}(ip)
	}

	wg.Wait()

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(*outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}
}

