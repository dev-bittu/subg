package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-bittu/subg/scanner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "subg",
  Short: "subg is small yet very powerful subdomain scanner",
  Long: `It is a very powerful subdomain scanner written in golang.`,
  Version: "0.0.1",
  Run: func(cmd *cobra.Command, args []string) {
    if Wordlist == "" || Domain == "" {
      fmt.Println("Enter Wordlist/Domain...")
      return
    }
    err := scanner.Scan(Domain, Wordlist, Thread)
    if err != nil {
      fmt.Println(err)
    }
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatalln(err)
    os.Exit(1)
  }
}
