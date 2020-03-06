package stream

import (
	"github.com/gocarina/gocsv"
	"github.com/kishaningithub/shopify-csv-download/internal/shopify"
	"io"
)

type ProductCSVStream struct {
	productCSV <-chan shopify.ProductCSV
}

func (stream ProductCSVStream) Save(out io.Writer) error {
	csvWriter := gocsv.DefaultCSVWriter(out)
	defer csvWriter.Flush()
	channel := stream.removeTypeFromProductCSVChannel(stream.productCSV)
	return gocsv.MarshalChan(channel, csvWriter)
}

func (stream ProductCSVStream) removeTypeFromProductCSVChannel(productCSV <-chan shopify.ProductCSV) <-chan interface{} {
	untypedProductCSV := make(chan interface{})
	go func() {
		defer close(untypedProductCSV)
		for csv := range productCSV {
			untypedProductCSV <- csv
		}
	}()
	return untypedProductCSV
}
