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

## Installation

Download the binary from releases from the browser or using curl

```bash
# For linux
curl -s https://api.github.com/repos/barelyhuman/dokcli/releases/latest | grep browser_download_url | grep linux-amd64 | cut -d '"' -f 4 | wget -qi -
tar -xvzf dokcli-linux-amd64.tar.gz
```

copy it to `/usr/local/bin` and make sure it's added to your PATH

```bash
cp dokcli /usr/local/bin
export PATH=$PATH:/usr/local/bin
```

## Usage

For now the cli just creates a script for you to setup dokku as needed.

1. Create a yml file named `dokku-gen.yml`. You can use the template provide [dokku-gen.template.yml](dokku-gen.template.yml)
2. Run `dokcli` in the needed folder
3. A script with the config's app name will be generated for you (screenshot below)

**Example of a generated script**

<img src="/.github/assets/script.png" height="400"  />

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
