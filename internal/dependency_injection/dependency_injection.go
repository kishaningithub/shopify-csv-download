package dependency_injection

import (
	"github.com/kishaningithub/shopify-csv-download/internal/resources"
	"github.com/kishaningithub/shopify-csv-download/internal/services"
	"net/url"
)

type RequiredObjects struct {
	ProductsCSVWriterService services.ProductsCSVWriterService
}

func ConstructRequiredObjects(shopifyStoreUrl url.URL) RequiredObjects{
	productCSVConversionService := services.NewProductCSVConversionService()
	shopifyResource := resources.NewShopifyResource(shopifyStoreUrl)
	productsRetrievalService := services.NewProductsRetrievalService(shopifyResource)
	productsCSVWriterService := services.NewProductsCSVWriterService(productCSVConversionService, productsRetrievalService)
	return RequiredObjects{
		ProductsCSVWriterService: productsCSVWriterService,
	}
}
