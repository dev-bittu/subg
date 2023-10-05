package subdomain

import (
	"os"
	"sync"
)

type Subdomains struct {
	subdomains []string
	mu         *sync.Mutex
}

func (s *Subdomains) AddSubdomain(subd string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subdomains = append(s.subdomains, subd)
}

func (s *Subdomains) WriteOnFile(file_path string) error {
	f, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, d := range s.subdomains {
		_, err := f.WriteString(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewSubdomains() *Subdomains {
	return &Subdomains{
		subdomains: []string{},
		mu:         &sync.Mutex{},
	}
}
