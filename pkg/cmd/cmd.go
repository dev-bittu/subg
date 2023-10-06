package cmd

import (
	"fmt"
	"os"

	alert "github.com/dev-bittu/goalert"
	"github.com/dev-bittu/subg/config"
	"github.com/dev-bittu/subg/internal/net"
	"github.com/dev-bittu/subg/pkg/scanner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     config.Config["name"].(string),
	Short:   config.Config["title"].(string),
	Long:    config.Config["desc"].(string),
	Version: config.Config["version"].(string),
	Example: config.Config["example"].(string),
	Run: func(cmd *cobra.Command, args []string) {
		if !net.IsOnline() {
			fmt.Println(alert.Red("You are not connected to internet,\nPlease check your wire."))
			os.Exit(0)
		}
		scanr, err := scanner.NewScanner(domain, wordlist, thread, outputFile, timeout)
		if err != nil {
			fmt.Println(err)
		}
		err = scanr.Scan()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(alert.Red("Solve the problem to move further:"))
		alert.Info("Problem: " + err.Error())
	}
}
