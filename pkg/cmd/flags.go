package cmd

var (
	Wordlist   string
	Domain     string
	Thread     int
	OutputFile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&Domain,
		"domain",
		"d",
		"",
		"Domain to scan. Example: example.com",
	)
	rootCmd.PersistentFlags().StringVarP(
		&Wordlist,
		"wordlist",
		"w",
		"",
		"Wordlist path (delimiter: newline or \\n). Example: wdlist/subdomains.txt",
	)
	rootCmd.PersistentFlags().StringVarP(
		&OutputFile,
		"output",
		"o",
		".subg.log",
		"Ouput file path (delimiter: newline or \\n). Example: subdomains.txt",
	)
	rootCmd.PersistentFlags().IntVarP(
		&Thread,
		"thread",
		"t",
		50,
		"Thread [subdomains/sec]",
	)

	rootCmd.MarkPersistentFlagRequired("domain")
	rootCmd.MarkPersistentFlagRequired("wordlist")
}
