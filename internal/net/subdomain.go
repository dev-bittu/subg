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

	// implement logic, how to find that subdomain exists

	if exists {
		s.writeOnFile(subdomain+"\n")
		fmt.Println(alert.Yellow("[+]"), alert.Green(subdomain))
	} else {
		fmt.Println(alert.Cyan("[-]"), alert.Red(subdomain))
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
