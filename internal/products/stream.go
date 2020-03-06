package products

import (
	"github.com/kishaningithub/shopify-csv-download/internal/products/stream"
	"github.com/kishaningithub/shopify-csv-download/internal/shopify"
	"net/url"
)

func Stream(productsResourceFullUrl url.URL) stream.ProductStream {
	resource := shopify.NewResource(productsResourceFullUrl)
	maxNoOfRecords := 250
	products := make(chan shopify.Product, maxNoOfRecords*3)
	errors := make(chan error)
	go func() {
		defer close(errors)
		defer close(products)
		pageNo := 1
		for {
			productResponse, err := resource.GetProducts(maxNoOfRecords, pageNo)
			if err != nil {
				errors <- err
				return
			}
			if len(productResponse.Products) == 0 {
				return
			}
			for _, product := range productResponse.Products {
				products <- product
			}
			pageNo++
		}
	}()
	return stream.ProductStream{
		Products: products,
		Errors:   errors,
	}
}
