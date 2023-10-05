package banner

import (
	"fmt"

	"github.com/dev-bittu/subg/config"
)

func ShowBanner() {
  text := fmt.Sprintf(`
#################################################

Tool:       subg      (github.com/dev-bittu/subg)
Written by: dev-bittu (github.com/dev-bittu)
Contact:    mail      (devbittu@proton.me)
Version:    %s     (check github for updates)

#################################################`, config.Config["version"])

  fmt.Println(text)
}
