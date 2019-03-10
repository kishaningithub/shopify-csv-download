# shopify csv download

[![Build Status](https://travis-ci.org/kishaningithub/shopify-csv-download.svg?branch=master)](https://travis-ci.org/kishaningithub/shopify-csv-download)
[![Go Doc](https://godoc.org/github.com/kishaningithub/shopify-csv-download?status.svg)](https://godoc.org/github.com/kishaningithub/shopify-csv-download)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/shopify-csv-download)](https://goreportcard.com/report/github.com/kishaningithub/shopify-csv-download)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Downloads](https://img.shields.io/github/downloads/kishaningithub/shopify-csv-download/latest/total.svg)](https://github.com/kishaningithub/shopify-csv-download/releases)
[![Latest release](https://img.shields.io/github/release/kishaningithub/shopify-csv-download.svg)](https://github.com/kishaningithub/shopify-csv-download/releases)

Download a shopify site in a csv format that the shopify importer understands

## Table of Contents

- [shopify csv download](#shopify-csv-download)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [Using Homebrew](#using-homebrew)
    - [Using Binary](#using-binary)
  - [Usage](#usage)
  - [Maintainers](#maintainers)
  - [Contribute](#contribute)
  - [License](#license)

## Install

### Using Homebrew

```bash
brew tap kishaningithub/tap
brew install shopify-csv-download
```

### Using Binary

```bash
# All unix environments with curl
curl -sfL https://raw.githubusercontent.com/kishaningithub/shopify-csv-download/master/install.sh | sudo sh -s -- -b /usr/local/bin

# In alpine linux (as it does not come with curl by default)
wget -O - -q https://raw.githubusercontent.com/kishaningithub/shopify-csv-download/master/install.sh | sudo sh -s -- -b /usr/local/bin
```

## Usage

```bash
shopify-csv-download https://shopify-site.com > shopify-site-products.csv
```

## Maintainers

[@kishaningithub](https://github.com/kishaningithub)

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2018 Kishan B
