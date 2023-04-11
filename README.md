# Burp Config Generator

A command-line tool to generate Burp Suite configuration files with custom IP scopes.

## Installation

To install `burp-config-generator`, make sure you have Go 1.16 or later installed, and run the following command:

```sh
go install github.com/<your-username>/burp-config-generator@latest
```

Make sure that your Go `bin` directory (usually `$GOPATH/bin` or `$HOME/go/bin`) is included in your system's `PATH` environment variable so you can use the installed binary directly.

## Usage

After installing `burp-config-generator`, you can generate a Burp Suite configuration file using the following command:

```sh
burp-config-generator --input path/to/scope.txt --output path/to/burp_config.json
```

By default, if the flags are not provided, the program will use "scope.txt" as the input file and "burp_config.json" as the output file:

```sh
burp-config-generator
```

The `scope.txt` file should contain a list of IP addresses, with one IP address per line. The generated configuration file will include these IP addresses as the in-scope targets for Burp Suite.
