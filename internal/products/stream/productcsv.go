package stream

import (
	"github.com/gocarina/gocsv"
	"github.com/kishaningithub/shopify-csv-download/internal/shopify"
	"io"
)

type ProductCSVStream struct {
	ProductCSV                 <-chan shopify.ProductCSV
	NoOfProductsConvertedAsCSV <-chan int
	ParentStream               ProductStream
}

func (stream ProductCSVStream) Save(out io.Writer) (ProgressState, <-chan error) {
	csvWriter := gocsv.DefaultCSVWriter(out)
	defer csvWriter.Flush()
	err := make(chan error)
	go func() {
		channel := stream.removeTypeFromProductCSVChannel(stream.ProductCSV)
		err <- gocsv.MarshalChan(channel, csvWriter)
	}()
	return ProgressState{
		NoOfProductsDownloaded:     stream.ParentStream.NoOfProductsDownloaded,
		NoOfProductsConvertedAsCSV: stream.NoOfProductsConvertedAsCSV,
	}, err
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
