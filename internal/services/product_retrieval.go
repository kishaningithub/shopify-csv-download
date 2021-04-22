package services

import (
	"github.com/kishaningithub/shopify-csv-download/internal/models/shopify"
	"github.com/kishaningithub/shopify-csv-download/internal/resources"
)

type ProductsRetrievalService interface {
	RetrieveAllProducts(chan<- shopify.Product) error
}

type productsRetrievalService struct {
	shopifyResource resources.ShopifyResource
}

func NewProductsRetrievalService(shopifyResource resources.ShopifyResource) ProductsRetrievalService {
	return &productsRetrievalService{
		shopifyResource: shopifyResource,
	}
}

func (service *productsRetrievalService) RetrieveAllProducts(products chan<- shopify.Product) error {
	maxRecordsPerPage := 250
	for pageNo := 1; ; pageNo++ {
		productsResponse, err := service.shopifyResource.GetProducts(maxRecordsPerPage, pageNo)
		if err != nil {
			close(products)
			return err
		}
		if len(productsResponse.Products) == 0 {
			close(products)
			return nil
		}
		for _, product := range productsResponse.Products {
			products <- product
		}
	}
}
