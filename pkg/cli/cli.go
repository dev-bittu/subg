package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-bittu/subg/config"
	"github.com/dev-bittu/subg/pkg/scanner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   config.Config["name"].(string),
  Short: config.Config["title"].(string),
  Long: config.Config["desc"].(string),
  Version: config.Config["version"].(string),
  Run: func(cmd *cobra.Command, args []string) {
    if Wordlist == "" || Domain == "" {
      fmt.Println("Enter Wordlist/Domain...")
      return
    }
    scanr, err := scanner.NewScanner(Domain, Wordlist, Thread)
    if err != nil {
      fmt.Println(err)
    }
    scanr.Scan()
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatalln(err)
    os.Exit(1)
  }
}
