# Senior Backend Technical Challenge

[![Go version](https://img.shields.io/badge/Go-go1.18-blue.svg)](https://go.dev/) [![tests](https://github.com/HectorMRC/backend-tech-test/actions/workflows/test.yaml/badge.svg?branch=master)](https://github.com/HectorMRC/backend-tech-test/actions/workflows/test.yaml)

## How to run

This project has been developed in such a way the application can be executed by command line, or using containers instead.

If you want to execute it directly on your terminal, it is as simple as executing the following command:
``` bash
$ make run
```

Otherwise, if you wish to run the application as a container, first you have to build it, and then you will be able to deploy it, as showing below:
``` bash
$ make build
$ make deploy
```
> Be aware that the container engines used for this project are [podman](https://podman.io/) and [podman-compose](https://github.com/containers/podman-compose). If you want to use another engine, maybe you will need to change some fields from the [compose](./compose.yaml) file and, for sure, the `build` and `deploy` commands from the [Makefile](./Makefile)

## Requirements

In order to work, the application has dependency with the three modules listed down below:

| Module | Version | Description |
|:--:|:--:|:--|
[go-chi](http://github.com/go-chi/chi) | v5.0.7 | Web router
[godotenv](http://github.com/joho/godotenv) | v1.4.0 | Environment variables loader
[zap](http://pkg.go.dev/go.uber.org/zap) | v5.0.7 | Blazing fast, structured, leveled logging in Go

