package services

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/kishaningithub/shopify-csv-download/internal/models/shopify"
	"github.com/kishaningithub/shopify-csv-download/pkg/progress"
	"golang.org/x/sync/errgroup"
	"io"
)

type ProductsCSVWriterService interface {
	DownloadAllProducts(out io.Writer) error
	DownloadAllProductsUpdatingProgressState(out io.Writer, progressState chan<- progress.State) error
}

type productsCSVWriterService struct {
	productCSVConversionService ProductCSVConversionService
	productsRetrievalService    ProductsRetrievalService
}

func NewProductsCSVWriterService(productCSVConversionService ProductCSVConversionService, productsRetrievalService ProductsRetrievalService) ProductsCSVWriterService {
	return &productsCSVWriterService{
		productCSVConversionService: productCSVConversionService,
		productsRetrievalService:    productsRetrievalService,
	}
}

func (service *productsCSVWriterService) DownloadAllProducts(out io.Writer) error {
	return service.downloadAllProductsUpdatingProgressStateIfRequired(out, nil)
}

func (service *productsCSVWriterService) DownloadAllProductsUpdatingProgressState(out io.Writer, progressState chan<- progress.State) error {
	return service.downloadAllProductsUpdatingProgressStateIfRequired(out, progressState)
}

func (service *productsCSVWriterService) downloadAllProductsUpdatingProgressStateIfRequired(out io.Writer, progressState chan<- progress.State) error {
	currentProgressState := progress.State{
		NoOfProductsDownloaded:     0,
		NoOfProductsConvertedAsCSV: 0,
	}
	csvWriter := make(chan interface{}, 1000)
	operation, _ := errgroup.WithContext(context.Background())
	operation.Go(func() error {
		return gocsv.MarshalChan(csvWriter, gocsv.DefaultCSVWriter(out))
	})
	products := make(chan shopify.Product, 1000)
	operation.Go(func() error {
		return service.productsRetrievalService.RetrieveAllProducts(products)
	})
	for product := range products {
		if progressState != nil {
			currentProgressState.NoOfProductsDownloaded++
			progressState <- currentProgressState
		}
		productCSVs := service.productCSVConversionService.ConvertToCSVFormat(product)
		if progressState != nil {
			currentProgressState.NoOfProductsConvertedAsCSV++
			progressState <- currentProgressState
		}
		for _, productCSV := range productCSVs {
			csvWriter <- productCSV
		}
	}
	close(csvWriter)
	return operation.Wait()
}
