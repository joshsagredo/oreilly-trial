# Oreilly Trial
[![CI](https://github.com/bilalcaliskan/oreilly-trial/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/oreilly-trial/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/oreilly-trial)](https://hub.docker.com/r/bilalcaliskan/oreilly-trial/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/oreilly-trial)](https://goreportcard.com/report/github.com/bilalcaliskan/oreilly-trial)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_oreilly-trial&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_oreilly-trial)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_oreilly-trial&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_oreilly-trial)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_oreilly-trial&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_oreilly-trial)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_oreilly-trial&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_oreilly-trial)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_oreilly-trial&metric=coverage)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_oreilly-trial)
[![Release](https://img.shields.io/github/release/bilalcaliskan/oreilly-trial.svg)](https://github.com/bilalcaliskan/oreilly-trial/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/oreilly-trial)](https://github.com/bilalcaliskan/oreilly-trial)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


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
  -h, --help              help for oreilly-trial
      --logLevel string   log level logging library (debug, info, warn, error) (default "info")
  -v, --version           version for oreilly-trial
```

> By default, oreilly-trial attempts to create trial account **--attemptCount** times. Default value of that flag is 10, if you can not create trial account in **--attemptCount** attempts, please increase that value in the range of 1-20.

## Installation

### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/oreilly-trial/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ ./oreilly-trial
```

### Homebrew
This project can be installed with [Homebrew](https://brew.sh/):
```shell
$ brew tap bilalcaliskan/tap
$ brew install bilalcaliskan/tap/oreilly-trial
```

Then similar to binary method, you can run it by calling below command:
```shell
$ oreilly-trial
```

### Docker
You can simply run docker image with default configuration:
```shell
$ docker run bilalcaliskan/oreilly-trial:latest
```

## Development
This project requires below tools while developing:
- [Golang 1.20](https://golang.org/doc/go1.20)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

Simply run below command to prepare your development environment:
```shell
$ python3 -m venv venv
$ source venv/bin/activate
$ pip3 install pre-commit
$ pre-commit install -c build/ci/.pre-commit-config.yaml
```

## License
Apache License 2.0
