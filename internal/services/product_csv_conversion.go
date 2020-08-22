package services

import (
	"github.com/kishaningithub/shopify-csv-download/internal/models/shopify"
	"strconv"
	"strings"
)

type ProductCSVConversionService interface {
	ConvertToCSVFormat(product shopify.Product) []shopify.ProductCSV
}

type productCSVConversionService struct {
}

func NewProductCSVConversionService() ProductCSVConversionService {
	return &productCSVConversionService{
	}
}

func (service *productCSVConversionService) ConvertToCSVFormat(product shopify.Product) []shopify.ProductCSV {
	var productsInCSV []shopify.ProductCSV
	noOfImagesInProduct := len(product.Images)
	for _, variant := range product.Variants {
		variantInventoryQuantity := "0"
		if variant.Available {
			variantInventoryQuantity = "1"
		}
		imageSrc := ""
		imagePosition := ""
		if noOfImagesInProduct > 0 {
			imageSrc = product.Images[0].Src
			imagePosition = strconv.Itoa(product.Images[0].Position)
		}
		option1Name := variant.Option1
		if variant.Option1 == "Default Title" {
			option1Name = "Title"
		}
		productsInCSV = append(productsInCSV, shopify.ProductCSV{
			Handle:                    product.Handle,
			Title:                     product.Title,
			BodyHTML:                  product.BodyHTML,
			Vendor:                    product.Vendor,
			Type:                      product.ProductType,
			Tags:                      strings.Join(product.Tags, ","),
			Published:                 strconv.FormatBool(true),
			Option1Name:               option1Name,
			Option1Value:              variant.Option1,
			Option2Name:               variant.Option2,
			Option2Value:              variant.Option2,
			Option3Name:               variant.Option3,
			Option3Value:              variant.Option3,
			VariantSKU:                variant.Sku,
			VariantGrams:              strconv.Itoa(variant.Grams),
			VariantInventoryTracker:   "shopify",
			VariantInventoryQty:       variantInventoryQuantity,
			VariantInventoryPolicy:    "deny",
			VariantFulfillmentService: "manual",
			VariantPrice:              variant.Price,
			VariantCompareAtPrice:     variant.CompareAtPrice,
			VariantRequiresShipping:   strconv.FormatBool(variant.RequiresShipping),
			VariantTaxable:            strconv.FormatBool(variant.Taxable),
			ImageSrc:                  imageSrc,
			ImagePosition:             imagePosition,
			GiftCard:                  strconv.FormatBool(false),
		})
	}
	if noOfImagesInProduct > 1 {
		for i := 1; i < noOfImagesInProduct; i++ {
			productImage := product.Images[i]
			imageSrc := productImage.Src
			imagePosition := strconv.Itoa(productImage.Position)
			productsInCSV = append(productsInCSV, shopify.ProductCSV{
				Handle:        product.Handle,
				ImageSrc:      imageSrc,
				ImagePosition: imagePosition,
			})
		}
	}
	return productsInCSV
}
