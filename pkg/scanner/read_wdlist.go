package scanner

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func getWordlist(file string) ([]string, error) {
  var wdlist []string
  f, err := os.Open(file)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  r := bufio.NewReader(f)

  for {
    line, err := r.ReadString('\n')
    if err == io.EOF {
      break
    } else if err != nil {
      return nil, err
    }
    line = strings.ReplaceAll(line, "\n", "")
    line = strings.ReplaceAll(line, "\r", "")

    wdlist = append(wdlist, line)
  }

  return wdlist, nil
}
