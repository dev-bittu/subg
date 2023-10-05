package cli

var (
  Wordlist string
  Domain string
  Thread uint
)

func init() {
  rootCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", "Domain to scan. Ex: example.com")
  rootCmd.PersistentFlags().StringVarP(&Wordlist, "wordlist", "w", "", "Wordlist path (delimiter: newline or \\n). Ex: /usr/share/wdlist/wd.txt")
  rootCmd.PersistentFlags().UintVarP(&Thread, "thread", "t", 100, "Thread (Default: 100) [subdomains/sec]")

  rootCmd.MarkFlagRequired("domain")
}
