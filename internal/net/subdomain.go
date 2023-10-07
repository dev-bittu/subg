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
	protocol   string
	Scan       uint16
}

func (s *Subdomains) incrementScan() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Scan++
}

func (s *Subdomains) Check(subdomain string) (bool, error) {
	s.incrementScan()
	exists := false
	url := fmt.Sprintf("%s://%s.%s", s.protocol, subdomain, s.domain)

	resp, err := http.Get(url)
	if err != nil {
		if err.Error() != fmt.Sprintf("Get \"%s\": dial tcp: lookup %s: no such host", url, subdomain+"."+s.domain) {
		return exists, err
	} else {
		return exists, nil
	}
	}
	if resp.StatusCode == 200 {
		exists = true
	}

	if exists {
		s.writeOnFile(subdomain + "\n")
		fmt.Println(
			alert.Green(fmt.Sprintf("[%d] %s", resp.StatusCode, subdomain)),
		)
	} else {
		fmt.Println(
			alert.Red(fmt.Sprintf("[%d] %s", resp.StatusCode, subdomain)),
		)
	}

	return exists, nil
}

func (s *Subdomains) writeOnFile(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.outputFile.WriteString(msg)
}

func (s *Subdomains) CloseFile() {
	s.outputFile.Close()
}

func NewSubdomains(domain string, timeout int, output string, protocol string) (*Subdomains, error) {
	f, err := os.Create(output)
	if err != nil {
		return nil, err
	}
	return &Subdomains{
		domain:     domain,
		subdomains: []string{},
		mu:         sync.Mutex{},
		outputFile: f,
		protocol:   protocol,
		client: http.Client{
			Timeout: time.Second * time.Duration(timeout),
		},
	}, nil
}
