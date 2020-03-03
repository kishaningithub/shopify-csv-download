package products

import (
	"github.com/kishaningithub/shopify-csv-download/internal/products"
	"io"
	"net/url"
)

func SaveAsImportableCSV(productsJsonURL url.URL, out io.Writer) error {
	return products.
		Stream(productsJsonURL).
		ConvertToCSV().
		Save(out)
}
