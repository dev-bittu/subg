package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/dev-bittu/subg/internal/net"
)

var (
	wg = sync.WaitGroup{}
	mu = sync.Mutex{}
)

func scanSubdomains(s *scanner, subdomain *net.Subdomains) {
	// Scan subdomains and write it to output file

	reader := bufio.NewReader(s.Wordlist)
	var (
		end     = false // indicate whether file is end or not
		count_t = 0     // count no of func executed concurrently
	)

	for !end {
		subd, err := reader.ReadString('\n')
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
			_ = subdomain.Check(subd + s.Domain)
		}(subd, err)
	}
	wg.Wait()
}
