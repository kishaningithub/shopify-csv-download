# shopify csv download

[![Build Status](https://github.com/kishaningithub/shopify-csv-download/actions/workflows/build.yml/badge.svg)](https://github.com/kishaningithub/shopify-csv-download/actions/workflows/build.yml)
[![Go Doc](https://godoc.org/github.com/kishaningithub/shopify-csv-download?status.svg)](https://godoc.org/github.com/kishaningithub/shopify-csv-download)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/shopify-csv-download)](https://goreportcard.com/report/github.com/kishaningithub/shopify-csv-download)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Latest release](https://img.shields.io/github/release/kishaningithub/shopify-csv-download.svg)](https://github.com/kishaningithub/shopify-csv-download/releases)
[![Buy me a lunch](https://img.shields.io/badge/🍱-Buy%20me%20a%20lunch-blue.svg)](https://www.paypal.me/kishansh/15)

Download a shopify site in a csv format that the [shopify importer understands](https://help.shopify.com/en/manual/products/import-export/using-csv#product-csv-file-format)

## Table of Contents

- [shopify csv download](#shopify-csv-download)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [Using Homebrew](#using-homebrew)
    - [Using Binary](#using-binary)
  - [Usage](#usage)
    - [CLI](#CLI)
    - [Library](#Library)
  - [Maintainers](#maintainers)
  - [Contribute](#contribute)
  - [License](#license)

## Install

### Using Homebrew

```bash
brew install kishaningithub/tap/shopify-csv-download
```

### Using Binary

```bash
# All unix environments with curl
curl -sfL https://raw.githubusercontent.com/kishaningithub/shopify-csv-download/master/install.sh | sudo sh -s -- -b /usr/local/bin

# In alpine linux (as it does not come with curl by default)
wget -O - -q https://raw.githubusercontent.com/kishaningithub/shopify-csv-download/master/install.sh | sudo sh -s -- -b /usr/local/bin
```

## Usage

### CLI

Retrieving all publicly exposed products

```bash
shopify-csv-download https://shopify-site.com > shopify-site-products.csv
```

### Library

```go
package main

import (
	"log"
	"net/url"
	"os"

	"github.com/kishaningithub/shopify-csv-download/pkg/products"
)

func main() {
	siteUrl, err := url.Parse("https://shopify-site.com")
	if err != nil {
		log.Println(err)
		return
	}
	err = products.SaveAsImportableCSV(*siteUrl, os.Stdout)
	if err != nil {
		log.Println(err)
		return
	}
}
```

## Maintainers

[@kishaningithub](https://github.com/kishaningithub)

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2021 Kishan B
