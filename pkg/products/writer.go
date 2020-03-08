package products

import (
	"github.com/kishaningithub/shopify-csv-download/internal/products"
	"github.com/kishaningithub/shopify-csv-download/internal/products/stream"
	"io"
	"net/url"
)

// SaveAsImportableCSV saves products from the given url to given writer
func SaveAsImportableCSV(productsJsonURL url.URL, out io.Writer) error {
	progressState, err := products.
		Stream(productsJsonURL).
		ConvertToCSV().
		Save(out)
	progressState.Ignore()
	return <-err
}

// SaveAsImportableCSVNotifyingProgressState saves products from the given url to given writer and notifies the progress state
func SaveAsImportableCSVNotifyingProgressState(productsJsonURL url.URL, out io.Writer, onProgressHandler ProgressHandler) error {
	progressState, err := products.
		Stream(productsJsonURL).
		ConvertToCSV().
		Save(out)
	onProgressChange(progressState, onProgressHandler)
	return <-err
}

func onProgressChange(progress stream.ProgressState, onProgressHandler ProgressHandler) {
	noOfProductsConvertedAsCSV := 0
	noOfProductsDownloaded := 0
	isNoOfProductsChannelOpen := false
	isNoOfProductsConvertedAsCSVChannelOpen := false
	for {
		select {
		case noOfProductsDownloaded, isNoOfProductsChannelOpen = <-progress.NoOfProductsDownloaded:
			if !isNoOfProductsChannelOpen {
				progress.NoOfProductsDownloaded = nil
				break
			}
			onProgressHandler(ProgressState{
				NoOfProductsDownloaded:     noOfProductsDownloaded,
				NoOfProductsConvertedAsCSV: noOfProductsConvertedAsCSV,
			})
		case noOfProductsConvertedAsCSV, isNoOfProductsConvertedAsCSVChannelOpen = <-progress.NoOfProductsConvertedAsCSV:
			if !isNoOfProductsConvertedAsCSVChannelOpen {
				progress.NoOfProductsConvertedAsCSV = nil
				break
			}
			onProgressHandler(ProgressState{
				NoOfProductsDownloaded:     noOfProductsDownloaded,
				NoOfProductsConvertedAsCSV: noOfProductsConvertedAsCSV,
			})
		}
		if progress.NoOfProductsDownloaded == nil && progress.NoOfProductsConvertedAsCSV == nil {
			break
		}
	}
}
