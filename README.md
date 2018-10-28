# shopify csv download

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

Download a shopify site in a csv format that the shopify importer understands

## Table of Contents

- [shopify csv download](#shopify-csv-download)
    - [Table of Contents](#table-of-contents)
    - [Install](#install)
    - [Usage](#usage)
    - [Maintainers](#maintainers)
    - [Contribute](#contribute)
    - [License](#license)

## Install

```bash
# binary will be downloaded to /usr/local/bin
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
