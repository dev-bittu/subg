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
	client     http.Client
	mu         sync.Mutex
}

func (s *Subdomains) AddSubdomain(subd string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subdomains = append(s.subdomains, subd)
}

func (s *Subdomains) Check(subdomain string) bool {
	exists := false

	if exists {
		s.AddSubdomain(subdomain + s.domain)
		fmt.Println(alert.Yellow("[+]"), alert.Green(subdomain))
	} else {
		fmt.Println(alert.Cyan("[-]"), alert.Red(subdomain))
	}

	return exists
}

func (s *Subdomains) WriteOnFile(file_path string) error {
	f, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, d := range s.subdomains {
		_, err := f.WriteString(d + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func NewSubdomains(domain string, timeout int) *Subdomains {
	return &Subdomains{
		domain:     domain,
		subdomains: []string{},
		mu:         sync.Mutex{},
		client: http.Client{
			Timeout: time.Second * time.Duration(timeout),
		},
	}
}
