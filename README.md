# Oreilly Trial

As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes.

This tool does couple of simple steps to provide free trial account for you:
  - Creates temp mail with specific domains over https://dropmail.p.rapidapi.com/
  - Tries to register with created temp mails to https://learning.oreilly.com/
  - Prints the login information to console and then exits.

## Configuration
oreilly-trial can be customized with several command line arguments:
```
Usage:
  oreilly-trial [flags]

Flags:
  -h, --help      help for oreilly-trial
      --verbose   verbose output of the logging library as 'debug' (default false)
  -v, --version   version for oreilly-trial
```

## Installation

### Binary
Binary can be downloaded from [Releases](https://github.com/joshsagredo/oreilly-trial/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ ./oreilly-trial
```

### Homebrew
This project can be installed with [Homebrew](https://brew.sh/):
```shell
$ brew tap joshsagredo/tap
$ brew install joshsagredo/tap/oreilly-trial
```

Then similar to binary method, you can run it by calling below command:
```shell
$ oreilly-trial
```

### Docker
You can simply run docker image with default configuration:
```shell
$ docker run joshsagredo/oreilly-trial:latest
```

## Development
This project requires below tools while developing:
- [Golang 1.20](https://golang.org/doc/go1.20)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/), simply run below command to prepare your development environment:
```shell
$ make pre-commit-setup
```

## License
Apache License 2.0
