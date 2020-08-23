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
	return &productCSVConversionService{}
}

func (service *productCSVConversionService) ConvertToCSVFormat(product shopify.Product) []shopify.ProductCSV {
	productsInCSV := make([]shopify.ProductCSV, 0, len(product.Variants))
	for _, variant := range product.Variants {
		productsInCSV = append(productsInCSV, service.getProductCSVForVariant(product, variant))
	}
	productsInCSV = append(productsInCSV, service.getProductCSVForAllImages(product)...)
	return productsInCSV
}

func (service *productCSVConversionService) getProductCSVForAllImages(product shopify.Product) []shopify.ProductCSV {
	noOfImagesInProduct := len(product.Images)
	productsInCSVForRemainingImages := make([]shopify.ProductCSV, 0, noOfImagesInProduct)
	for _, image := range product.Images {
		productsInCSVForRemainingImages = append(productsInCSVForRemainingImages, service.getProductCSVForImage(product, image))
	}
	return productsInCSVForRemainingImages
}

func (service *productCSVConversionService) getProductCSVForImage(product shopify.Product, image shopify.Image) shopify.ProductCSV {
	return shopify.ProductCSV{
		Handle:        product.Handle,
		ImageSrc:      image.Src,
		ImagePosition: strconv.Itoa(image.Position),
	}
}

func (service *productCSVConversionService) getProductCSVForVariant(product shopify.Product, variant shopify.Variant) shopify.ProductCSV {
	return shopify.ProductCSV{
		Handle:                    product.Handle,
		Title:                     product.Title,
		BodyHTML:                  product.BodyHTML,
		Vendor:                    product.Vendor,
		Type:                      product.ProductType,
		Tags:                      strings.Join(product.Tags, ","),
		Published:                 strings.ToUpper(strconv.FormatBool(true)),
		Option1Name:               service.getOption1Name(product, variant),
		Option1Value:              variant.Option1,
		Option2Name:               service.getOption2Name(product, variant),
		Option2Value:              variant.Option2,
		Option3Name:               service.getOption3Name(product, variant),
		Option3Value:              variant.Option3,
		VariantSKU:                variant.Sku,
		VariantGrams:              strconv.Itoa(variant.Grams),
		VariantInventoryTracker:   "shopify",
		VariantInventoryQty:       service.getVariantInventoryQuantity(variant),
		VariantInventoryPolicy:    "deny",
		VariantFulfillmentService: "manual",
		VariantPrice:              variant.Price,
		VariantCompareAtPrice:     variant.CompareAtPrice,
		VariantBarcode:            variant.Barcode,
		VariantRequiresShipping:   strings.ToUpper(strconv.FormatBool(variant.RequiresShipping)),
		VariantTaxable:            strings.ToUpper(strconv.FormatBool(variant.Taxable)),
		GiftCard:                  strings.ToUpper(strconv.FormatBool(false)),
	}
}

func (service *productCSVConversionService) getOption1Name(product shopify.Product, variant shopify.Variant) string {
	return product.Options[0].Name
}

func (service *productCSVConversionService) getOption2Name(product shopify.Product, variant shopify.Variant) string {
	if len(variant.Option2) > 0 {
		return product.Options[1].Name
	}
	return ""
}

func (service *productCSVConversionService) getOption3Name(product shopify.Product, variant shopify.Variant) string {
	if len(variant.Option3) > 0 {
		return product.Options[2].Name
	}
	return ""
}

func (service *productCSVConversionService) getVariantInventoryQuantity(variant shopify.Variant) string {
	variantInventoryQuantity := "0"
	if variant.Available {
		variantInventoryQuantity = "1"
	}
	return variantInventoryQuantity
}
