package scanner

import (
	"errors"
	"os"

	"github.com/dev-bittu/subg/internal/banner"
	"github.com/dev-bittu/subg/internal/file"
	"github.com/dev-bittu/subg/pkg/subdomain"
)

type scanner struct {
	Domain       string
	WordlistPath string
	Thread       int
	OutputFile   string
	Wdlst        *os.File
}

func NewScanner(domain string, wdlist string, thread int) (*scanner, error) {
	f, err := os.Open(wdlist)
	if err != nil {
		return nil, err
	}

	return &scanner{
		Domain:       domain,
		WordlistPath: wdlist,
		Thread:       thread,
		OutputFile:   ".subg_logs",
		Wdlst:        f,
	}, nil
}

func (s *scanner) Scan() error {
	defer s.Wdlst.Close()

	banner.ShowBanner(s.Domain, s.Thread)

	fe, err := file.IsFileExists(s.WordlistPath)
	if err != nil {
		return err
	}
	if !fe {
		return errors.New("Wordlist " + s.WordlistPath + " not exists")
	}

	sbd := subdomain.NewSubdomains()

	scanSubdomains(*s, sbd)

	err = sbd.WriteOnFile(s.OutputFile)
	if err != nil {
		return err
	}

	return nil
}
