package banner

import (
	"fmt"

	"github.com/dev-bittu/subg/config"
)

func ShowBanner(domain string, thread int) {
	text := fmt.Sprintf(`
#################################################

Tool:       subg      (github.com/dev-bittu/subg)
Written by: dev-bittu (github.com/dev-bittu)
Contact:    mail      (devbittu@proton.me)

Version:    %s

Domain:     %s
Thread:     %d

#################################################`,
		config.Config["version"], domain, thread,
	)

	fmt.Println(text)
}
