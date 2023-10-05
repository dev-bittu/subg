package scanner

import (
	"errors"
	"fmt"
	//"time"
)

func Scan(domain string, wd_file string, thread uint) error {
  fmt.Println("Start Scanning")

  fmt.Println("Checking if wdlist exists")
  fe, err := isFileExists(wd_file)
  if err != nil {
    return err
  }
  if !fe {
    return errors.New("Wordlist "+wd_file+" not exists")
  }

  fmt.Println("Extracting words from wordlist")
  wdlst, err := getWordlist(wd_file)
  if err != nil {
    return err
  }

  fmt.Println("Extracting subdomains")
  subds := getAllSubd(wdlst, domain, thread)

  fmt.Println(subds)
  // fmt.Println("\nWriting all sibdomains to file")
 
  return nil
}

func getAllSubd(wdlst []string, domain string, thread uint) []string {
  var (
    activeSub = []string{}
    sendReq uint = 0
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
