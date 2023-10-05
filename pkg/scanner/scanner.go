package scanner

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/dev-bittu/subg/internal/banner"
	"github.com/dev-bittu/subg/internal/file"
	//"time"
)

type scanner struct {
  Domain string
  WordlistPath string
  Thread int
  OutputFile string
  wg *sync.WaitGroup
  Wdlst *os.File
}

func NewScanner(domain string, wdlist string, thread int) (*scanner, error){
  f, err := os.Open(wdlist)
  if err != nil {
    return nil, err
  }

  return &scanner{
    Domain: domain,
    WordlistPath: wdlist,
    Thread: thread,
    OutputFile: ".subg_logs",
    Wdlst: f,
    wg: &sync.WaitGroup{},
  }, nil
}

func (s *scanner) Scan() error {
  banner.ShowBanner()

  fmt.Println("Checking if wdlist exists")
  fe, err := file.IsFileExists(s.WordlistPath)
  if err != nil {
    return err
  }
  if !fe {
    return errors.New("Wordlist "+s.WordlistPath+" not exists")
  }

  fmt.Println("Loading wordlist")
  wdlst, err := getWordlist(s.WordlistPath)
  if err != nil {
    return err
  }

  fmt.Println("Brute Force..")
  subds := getAllSubd(wdlst, s.Domain, s.Thread)

  fmt.Println(subds)
  // fmt.Println("\nWriting all sibdomains to file")
 
  return nil
}

func getAllSubd(wdlst []string, domain string, thread int) []string {
  var (
    activeSub = []string{}
    sendReq int = 0
  )
  for _, wd := range wdlst {
    if sendReq == thread {
      sendReq = 0
      //time.Sleep(time.Second*1)
    }
    e := isSubdExists(wd+"."+domain)
    if e {
      activeSub = append(activeSub, wd)
      fmt.Println(wd)
    }
  }

  return activeSub
}
