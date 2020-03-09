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
}

func main() {
	productsJsonURL, err := url.Parse(parseArgsAndGetStoreUrl())
	exitOnFailure(fmt.Sprintf("unable to parse url %s", productsJsonURL), err)
	logWithNewLine("Downloading products as CSV...")
	watch := stopwatch.Start()
	err = products.SaveAsImportableCSVNotifyingProgressState(*productsJsonURL, os.Stdout, progressHandler)
	exitOnFailure("unable to write products", err)
	logWithNewLine("")
	watch.Stop()
	logWithNewLine("Save complete. Time taken %s", watch.String())
}

func progressHandler(state products.ProgressState) {
	progressStateLineFormat := "%d products downloaded..."
	logInTheSameLine(progressStateLineFormat, state.NoOfProductsConvertedAsCSV)
}

func logWithNewLine(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	_, _ = fmt.Fprintln(os.Stderr)
}

func logInTheSameLine(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "\r"+format, args...)
}

func parseArgsAndGetStoreUrl() string {
	remainingArgs, err := flags.Parse(&opts)
	exitOnFailure("unable to parse flags", err)
	baseURL := strings.Join(remainingArgs, "")
	return baseURL
}

func exitOnFailure(message string, err error) {
	if err != nil {
		logWithNewLine(fmt.Errorf("%s: %w", message, err).Error())
		os.Exit(1)
	}
}
