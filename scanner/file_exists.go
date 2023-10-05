package scanner

import (
	"errors"
	"os"
)

func isFileExists(file string) (bool, error) {
  _, err := os.Stat(file)
  if err == nil {
    return true, nil
  } else if errors.Is(err, os.ErrNotExist) {
    return false, nil
  } 
  return false, err
}
