# dokcli

[![Binary Builds](https://github.com/barelyhuman/dokcli/actions/workflows/create-binaries.yml/badge.svg)](https://github.com/barelyhuman/dokcli/actions/workflows/create-binaries.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/barelyhuman/dokcli)](https://goreportcard.com/report/github.com/barelyhuman/dokcli)

## Motivation

I use [dokku](https://github.com/dokku/dokku) a for most deployments and while their cli is pretty extensive a certain set of operations could be better in terms of creating an App.

There's always predefined steps for various apps and hence I'd like to have a cli that can act as a wrapper around dokku to make it easier for me to create apps.

## Key Features (For now)

- [x] Create a New App
  - [x] Enter app name
  - [x] Select database plugin
  - [x] link database to app
  - [x] Add domain to the app
  - [x] Add Let's encrypt
- [ ] Delete App
  - [ ] Unlink all the above
  - [ ] clean up un-used images

## Roadmap

- Cli for creation and deletion of app
- Lightweight Web-UI for the same

## Installation

Download the binary from releases from the browser or using curl

```bash
# For linux / unix, make sure wget is installed
curl -s https://api.github.com/repos/barelyhuman/dokcli/releases/latest | grep browser_download_url | grep linux-amd64 | cut -d '"' -f 4 | wget -qi -
tar -xvzf dokcli-linux-amd64.tar.gz
```

copy it to `/usr/local/bin` and make sure it's added to your PATH

```bash
cp dokcli /usr/local/bin
export PATH=$PATH:/usr/local/bin
```

## Usage

**Note: Make sure you have dokku installed**

For now the cli just creates a script for you to setup a dokku as needed.

(**Step 0 is optional**), the cli will ask you the needed information.

0. Create a yml file named `dokku-gen.yml`. You can use the provided template  [dokku-gen.template.yml](dokku-gen.template.yml)

1. Run `dokcli` on the system to create the app on.
2. A script with the app name will be generated for you (screenshot below) with the name of `dokku-setup-<app-name>.sh`
3. `chmod +x ./dokku-setup-<app-name>.sh`
4. `./dokku-setup-<app-name>.sh` to run the generated script.

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
