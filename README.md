## Oreilly Trial

[![CI](https://github.com/bilalcaliskan/oreilly-trial/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/oreilly-trial/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/oreilly-trial)](https://hub.docker.com/r/bilalcaliskan/oreilly-trial/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/oreilly-trial)](https://goreportcard.com/report/github.com/bilalcaliskan/oreilly-trial)

As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes. 

This tool does couple of simple steps to provide free trial account for you:
  - Register with temp mail to https://learning.oreilly.com/
  - Print the login information to console and then exit.

### Configuration
oreilly-trial can be customized with several command line arguments:
```
--createUserUrl             url of the user creation on Oreilly API. Defaults to https://learning.oreilly.com/api/v1/user/
--emailDomains              comma seperated list of usable domain for creating trial account, it should be a valid domain. Defaults to {"jentrix.com", "geekale.com", "64ge.com", "frnla.com"}
--length                    length of the random generated username and password. Defaults to 16.
```

### Download

#### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/oreilly-trial/releases) page.

#### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/oreilly-trial:latest
```
_**Happy learning!**_
