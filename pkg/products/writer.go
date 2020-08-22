package products

import (
	"fmt"
	"github.com/kishaningithub/shopify-csv-download/internal/dependency_injection"
	"github.com/kishaningithub/shopify-csv-download/pkg/progress"
	"io"
	"net/url"
)

// SaveAsImportableCSV saves products from the given url to given writer
func SaveAsImportableCSV(shopifyStoreUrlString string, out io.Writer) error {
	shopifyStoreUrl, err := url.ParseRequestURI(shopifyStoreUrlString)
	if err != nil {
		return fmt.Errorf("invalid url %s: %w", shopifyStoreUrlString, err)
	}
	productCSVWriter := dependency_injection.ConstructRequiredObjects(*shopifyStoreUrl).ProductsCSVWriterService
	return productCSVWriter.DownloadAllProducts(out)
}

// SaveAsImportableCSVNotifyingProgressState saves products from the given url to given writer and notifies the progress state
func SaveAsImportableCSVNotifyingProgressState(shopifyStoreUrl url.URL, out io.Writer, onProgressHandler progress.Handler) error {
	progressStates := make(chan progress.State, 1000)
	go func() {
		for progressState := range progressStates {
			onProgressHandler(progressState)
		}
	}()
	productCSVWriter := dependency_injection.ConstructRequiredObjects(shopifyStoreUrl).ProductsCSVWriterService
	return productCSVWriter.DownloadAllProductsUpdatingProgressState(out, progressStates)
}