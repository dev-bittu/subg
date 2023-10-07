package net

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	alert "github.com/dev-bittu/goalert"
)

type Subdomains struct {
	domain     string
	subdomains []string
	outputFile *os.File
	client     http.Client
	mu         sync.Mutex
}

func (s *Subdomains) Check(subdomain string) bool {
	exists := false

	resp, err := http.Get(subdomain + "." + s.domain)
	if err != nil {
		return exists
	} 
	if resp.StatusCode == 200 {
		exists = true
	}

	if exists {
		s.writeOnFile(subdomain + "\n")
		fmt.Println(
			alert.Yellow(fmt.Sprintf("[%d] %s", resp.StatusCode, subdomain)),
		)
	} else {
		fmt.Println(
    	alert.Green(fmt.Sprintf("[%d] %s", resp.StatusCode, subdomain)),
    )
	}

	return exists
}

func (s *Subdomains) writeOnFile(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.outputFile.WriteString(msg)
}

func (s *Subdomains) CloseFile() {
	s.outputFile.Close()
}

func NewSubdomains(domain string, timeout int, output string) (*Subdomains, error) {
	f, err := os.Create(output)
	if err != nil {
		return nil, err
	}
	return &Subdomains{
		domain:     domain,
		subdomains: []string{},
		mu:         sync.Mutex{},
		outputFile: f,
		client: http.Client{
			Timeout: time.Second * time.Duration(timeout),
		},
	}, nil
}
