package main

import (
	"fmt"
	"os"

	alert "github.com/dev-bittu/goalert"
	"github.com/dev-bittu/subg/internal/net"
	"github.com/dev-bittu/subg/pkg/cmd"
)

func init() {
	if !net.IsOnline() {
		fmt.Println(alert.Red("You are not connected to internet,\nPlease check your wire."))
		os.Exit(0)
	}
}

func main() {
	cmd.Execute()
}
