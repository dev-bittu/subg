package scanner

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/dev-bittu/subg/pkg/subdomain"
)

var (
	wg = sync.WaitGroup{}
	mu = sync.Mutex{}
)

func scanSubdomains(s scanner, subdomain *subdomain.Subdomains) {
	// Scan subdomains and write it to output file

	reader := bufio.NewReader(s.Wdlst)
	var (
		end     = false // indicate whether file is end or not
		count_t = 0     // count no of func executed concurrently
	)

	for !end {
		sub, err := reader.ReadString('\n')
		if count_t == s.Thread {
			count_t = 0
			time.Sleep(time.Second * 1)
		}

		wg.Add(1)
		go func(subd string, err error) {
			defer wg.Done()
			if err == io.EOF {
				end = true
				return
			} else if err != nil {
				fmt.Println(err)
				return
			}

			mu.Lock()
			count_t++
			mu.Unlock()

			subd = strings.ReplaceAll(
				strings.ReplaceAll(subd, "\n", ""),
				"\r",
				"",
			)
			exists := isSubdExists(subd + s.Domain)
			if exists {
				fmt.Println(subd)
				subdomain.AddSubdomain(subd)
			}
		}(sub, err)
	}
	wg.Wait()
}

func isSubdExists(d string) bool {
	// Check if given subdomain exists
	_, err := net.LookupHost(d)
	return err == nil
}
