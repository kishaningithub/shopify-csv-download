package main

import (
	"fmt"
	"github.com/bradhe/stopwatch"
	"github.com/jessevdk/go-flags"
	"github.com/kishaningithub/shopify-csv-download/pkg/products"
	"net/url"
	"os"
	"strings"
)

var opts struct {
	FullUrl string `short:"f" long:"full-url" description:"Full URL to products.json. Eg: https://shopify-site.com/products.json"`
}

func main() {
	productsJsonURL, err := url.Parse(findProductsJsonURL())
	exitOnFailure(fmt.Sprintf("unable to parse url %s", productsJsonURL), err)
	logWithNewLine("Downloading products as CSV...")
	watch := stopwatch.Start()
	err = products.SaveAsImportableCSVWithProgressState(*productsJsonURL, os.Stdout, progressHandler)
	exitOnFailure("unable to write products", err)
	logWithNewLine("")
	watch.Stop()
	logWithNewLine("Save complete. Time taken %s", watch.String())
}

func progressHandler(state products.ProgressState) {
	progressStateLineFormat := "Products downloaded: %d Products converted as CSV: %d"
	logInTheSameLine(progressStateLineFormat, state.NoOfProductsDownloaded, state.NoOfProductsConvertedAsCSV)
}

func logWithNewLine(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	_, _ = fmt.Fprintln(os.Stderr)
}

func logInTheSameLine(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "\r"+format, args...)
}

func findProductsJsonURL() string {
	remainingArgs, err := flags.Parse(&opts)
	exitOnFailure("unable to parse flags", err)
	baseURL := strings.Join(remainingArgs, "")
	productsJsonURL := fmt.Sprintf("%s/products.json", baseURL)
	if len(opts.FullUrl) > 0 {
		productsJsonURL = opts.FullUrl
	}
	logWithNewLine("Products URL is %s", productsJsonURL)
	return productsJsonURL
}

func exitOnFailure(message string, err error) {
	if err != nil {
		logWithNewLine(fmt.Errorf("%s: %w", message, err).Error())
		os.Exit(1)
	}
}
