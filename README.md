# SUBG - SUBdomain scanner written in Golang
Subg is a subdomain scanner written in Golang.
It is a command-line tool that allows you to scan a domain for subdomains.

## Installation
### Simple Approach
Install subg using go (golang):
```bash
go install github.com/dev-bittu/subg
```
> **NOTE:** Make sure go/bin is in your PATH

### Manual Approach
1. Clone this repo:
```bash
git clone https://github.com/dev-bittu/subg
```
2. Change directory:
```bash
cd subg
```
3. Compile subg:
```bash
go build .
```
Now, you can simply check help by ./sibg -h

4. Add to path:
```bash
cp subg $PATH/subg
```

## Usage
To use subg,
simply run the following command:
```bash
subg -d example.com -w <you_wdlst>
```
This will scan the example.com domain for subdomains.

### Flags
Subg has the following flags:
- --domain or -d: The domain to scan (required)
- --thread or -t: The number of threads to use (default: 100)
- --wordlist or -w: The path to the wordlist file (required)

### Examples
Scan a domain with default settings:
```bash
subg -d example.com
```

Scan a domain with 50 threads:
```bash
subg -d example.com -t 50
```

Scan a domain with a custom wordlist:
```bash
subg -t example.com -w /path/to/custom/wordlist.txt
```

## Contributing
If you'd like to contribute to subg, please create a pull request with your changes.
Before submitting a pull request, please make sure that your changes pass the tests and that they adhere to the project's coding standards.

## License
Subg is licensed under the MIT license.
See the [License](LICENSE) file for more information.
