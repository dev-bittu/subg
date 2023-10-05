package scanner

import "net"

func isSubdExists(d string) bool {
  _, err := net.LookupHost(d)
  return err == nil
}
