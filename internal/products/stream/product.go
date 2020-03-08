package stream

import (
	"github.com/kishaningithub/shopify-csv-download/internal/products/convertor"
	"github.com/kishaningithub/shopify-csv-download/internal/shopify"
)

type ProductStream struct {
	Products               <-chan shopify.Product
	Errors                 <-chan error
	NoOfProductsDownloaded <-chan int
}

func (stream ProductStream) ConvertToCSV() ProductCSVStream {
	productCSV := make(chan shopify.ProductCSV)
	noOfProductsConvertedAsCSV := make(chan int)
	go func() {
		defer close(productCSV)
		defer close(noOfProductsConvertedAsCSV)
		noOfProductsConvertedAsCSVCounter := 0
		for product := range stream.Products {
			for _, csv := range convertor.ConvertToCSVFormat(product) {
				productCSV <- csv
			}
			noOfProductsConvertedAsCSVCounter += 1
			noOfProductsConvertedAsCSV <- noOfProductsConvertedAsCSVCounter
		}
	}()
	return ProductCSVStream{
		ParentStream:               stream,
		ProductCSV:                 productCSV,
		NoOfProductsConvertedAsCSV: noOfProductsConvertedAsCSV,
	}
}
