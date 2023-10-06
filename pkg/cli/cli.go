package cli

import (
	"fmt"
	"net"
	"net/http"
	"time"

	alert "github.com/dev-bittu/go-alert"
	"github.com/dev-bittu/subg/config"
	"github.com/dev-bittu/subg/pkg/scanner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     config.Config["name"].(string),
	Short:   config.Config["title"].(string),
	Long:    config.Config["desc"].(string),
	Version: config.Config["version"].(string),
	Run: func(cmd *cobra.Command, args []string) {
		if !isConnectedToInternet() {
			fmt.Println(alert.Red())
			return
		}
		scanr, err := scanner.NewScanner(Domain, Wordlist, Thread)
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
		fmt.Println("Solve the problem to move further:")
		fmt.Println("Problem:", err)
	}
}

func isConnectedToInternet() bool {
	// Create a timeout for the request
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Make a GET request to a reliable server
	resp, err := client.Get("https://www.example.com")
	if err != nil {
		// Check if the error is due to no internet connection
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return false
		}
		return false
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	}
	return false
}
