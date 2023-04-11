# Burp Config Generator

A command-line tool to generate Burp Suite configuration files with custom IP scopes.

## Installation

To install `burpconfiggenerator`, make sure you have Go 1.16 or later installed, and run the following command:

```sh
go install github.com/predator77799/burpconfiggenerator@latest
```

Make sure that your Go `bin` directory (usually `$GOPATH/bin` or `$HOME/go/bin`) is included in your system's `PATH` environment variable so you can use the installed binary directly.

## Usage

After installing `burpconfiggenerator`, you can generate a Burp Suite configuration file using the following command:

```sh
burpconfiggenerator --input path/to/scope.txt --output path/to/burp_config.json
```

By default, if the flags are not provided, the program will use "scope.txt" as the input file and "burp_config.json" as the output file:

```sh
burpconfiggenerator
```

The `scope.txt` file should contain a list of IP addresses, with one IP address per line. The generated configuration file will include these IP addresses as the in-scope targets for Burp Suite.
