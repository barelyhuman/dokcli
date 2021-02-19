# dokcli

[![Binary Builds](https://github.com/barelyhuman/dokcli/actions/workflows/create-binaries.yml/badge.svg)](https://github.com/barelyhuman/dokcli/actions/workflows/create-binaries.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/barelyhuman/dokcli)](https://goreportcard.com/report/github.com/barelyhuman/dokcli)

## Motivation

I use [dokku](https://github.com/dokku/dokku) a for most deployments and while their cli is pretty extensive a certain set of operations could be better in terms of creating an App.

There's always predefined steps for various apps and hence I'd like to have a cli that can act as a wrapper around dokku to make it easier for me to create apps.

## Key Features (For now)

- [ ] Create a New App
  - [ ] Enter app name
  - [ ] Select database plugin
  - [ ] link database to app
  - [ ] Add domain to the app
  - [ ] Add Let's encrypt
- [ ] Delete App
  - [ ] Unlink all the above
  - [ ] clean up un-used images

## Roadmap

- Cli for creation and deletion of app
- Lightweight Web-UI for the same

## Dev

Make sure you have a minimum of `go 1.5`
_Run_

```bash
    go run .
```

_Build_

```bash
go build .
```

## Contribution

Check the issues for things you can pick up and send through a PR for
