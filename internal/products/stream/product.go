package stream

import (
	"github.com/kishaningithub/shopify-csv-download/internal/products/convertor"
	"github.com/kishaningithub/shopify-csv-download/internal/shopify"
)

type ProductStream struct {
	Products <-chan shopify.Product
	Errors   <-chan error
}

func (stream ProductStream) ConvertToCSV() ProductCSVStream {
	productCSV := make(chan shopify.ProductCSV)
	go func() {
		defer close(productCSV)
		for product := range stream.Products {
			for _, csv := range convertor.ConvertToCSVFormat(product) {
				productCSV <- csv
			}
		}
	}()
	return ProductCSVStream{
		productCSV: productCSV,
	}
}
