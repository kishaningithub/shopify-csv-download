package services_test

import (
	"github.com/kishaningithub/shopify-csv-download/internal/models/shopify"
	"github.com/kishaningithub/shopify-csv-download/internal/services"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	_ suite.SetupTestSuite = (*ProductCSVConversionServiceTestSuite)(nil)
)

type ProductCSVConversionServiceTestSuite struct {
	suite.Suite
	productCSVConversionService services.ProductCSVConversionService
}

func TestProductCSVConversionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductCSVConversionServiceTestSuite))
}

func (suite *ProductCSVConversionServiceTestSuite) SetupTest() {
	suite.productCSVConversionService = services.NewProductCSVConversionService()
}

func (suite *ProductCSVConversionServiceTestSuite) TestProductConversionForAProductWithAllInformation() {
	productWithAllInformation := shopify.Product{
		Title:       "title",
		Handle:      "handle",
		BodyHTML:    "bodyHTML",
		Vendor:      "vendor",
		ProductType: "productType",
		Tags:        []string{"tag1", "tag2"},
		Variants: []shopify.Variant{
			{
				Option1:          "option1Value1",
				Option2:          "option2Value1",
				Option3:          "option3Value1",
				Sku:              "variant1Sku",
				RequiresShipping: true,
				Taxable:          true,
				Barcode:          "variant1Barcode",
				Available:        true,
				Price:            "1000",
				Grams:            100,
				CompareAtPrice:   "10000",
			},
			{
				Option1:          "option1Value2",
				Option2:          "option2Value2",
				Option3:          "option3Value2",
				Sku:              "variant2Sku",
				RequiresShipping: true,
				Taxable:          true,
				Barcode:          "variant2Barcode",
				Available:        true,
				Price:            "2000",
				Grams:            200,
				CompareAtPrice:   "20000",
			},
		},
		Images: []shopify.Image{
			{
				Position: 1,
				Src:      "image1Src",
			},
			{
				Position: 2,
				Src:      "image2Src",
			},
			{
				Position: 3,
				Src:      "image3Src",
			},
		},
		Options: []shopify.Option{
			{
				Name:     "Option 1 name",
				Position: 1,
				Values: []string{
					"option1Value1",
					"option1Value2",
					"option1Value3",
				},
			},
			{
				Name:     "Option 2 name",
				Position: 2,
				Values: []string{
					"option2Value1",
					"option2Value2",
					"option2Value3",
				},
			},
			{
				Name:     "Option 3 name",
				Position: 3,
				Values: []string{
					"option3Value1",
					"option3Value2",
					"option3Value3",
				},
			},
		},
	}

	expectedProductCSV := []shopify.ProductCSV{
		// Variant 1
		{
			Handle:                    "handle",
			Title:                     "title",
			BodyHTML:                  "bodyHTML",
			Vendor:                    "vendor",
			Type:                      "productType",
			Tags:                      "tag1,tag2",
			Published:                 "TRUE",
			Option1Name:               "Option 1 name",
			Option1Value:              "option1Value1",
			Option2Name:               "Option 2 name",
			Option2Value:              "option2Value1",
			Option3Name:               "Option 3 name",
			Option3Value:              "option3Value1",
			VariantSKU:                "variant1Sku",
			VariantGrams:              "100",
			VariantInventoryTracker:   "shopify",
			VariantInventoryQty:       "1",
			VariantInventoryPolicy:    "deny",
			VariantFulfillmentService: "manual",
			VariantPrice:              "1000",
			VariantCompareAtPrice:     "10000",
			VariantRequiresShipping:   "TRUE",
			VariantTaxable:            "TRUE",
			VariantBarcode:            "variant1Barcode",
			GiftCard:                  "FALSE",
		},
		// Variant 2
		{
			Handle:                    "handle",
			Title:                     "title",
			BodyHTML:                  "bodyHTML",
			Vendor:                    "vendor",
			Type:                      "productType",
			Tags:                      "tag1,tag2",
			Published:                 "TRUE",
			Option1Name:               "Option 1 name",
			Option1Value:              "option1Value2",
			Option2Name:               "Option 2 name",
			Option2Value:              "option2Value2",
			Option3Name:               "Option 3 name",
			Option3Value:              "option3Value2",
			VariantSKU:                "variant2Sku",
			VariantGrams:              "200",
			VariantInventoryTracker:   "shopify",
			VariantInventoryQty:       "1",
			VariantInventoryPolicy:    "deny",
			VariantFulfillmentService: "manual",
			VariantPrice:              "2000",
			VariantCompareAtPrice:     "20000",
			VariantRequiresShipping:   "TRUE",
			VariantTaxable:            "TRUE",
			VariantBarcode:            "variant2Barcode",
			GiftCard:                  "FALSE",
		},
		// Images
		{
			Handle:        "handle",
			ImageSrc:      "image1Src",
			ImagePosition: "1",
		},
		{
			Handle:        "handle",
			ImageSrc:      "image2Src",
			ImagePosition: "2",
		},
		{
			Handle:        "handle",
			ImageSrc:      "image3Src",
			ImagePosition: "3",
		},
	}

	actualProductCSV := suite.productCSVConversionService.ConvertToCSVFormat(productWithAllInformation)

	suite.Require().Equal(expectedProductCSV, actualProductCSV)
}

func (suite *ProductCSVConversionServiceTestSuite) TestProductConversionForAProductWithNoVariants() {
	productWithAllInformation := shopify.Product{
		Title:       "title",
		Handle:      "handle",
		BodyHTML:    "bodyHTML",
		Vendor:      "vendor",
		ProductType: "productType",
		Tags:        []string{"tag1", "tag2"},
		Variants: []shopify.Variant{
			{
				Title:            "variant1Title",
				Option1:          "Default Title",
				Sku:              "sku",
				RequiresShipping: true,
				Taxable:          true,
				Barcode:          "barcode",
				Available:        true,
				Price:            "1000",
				Grams:            100,
				CompareAtPrice:   "10000",
			},
		},
		Options: []shopify.Option{
			{
				Name:     "Title",
				Position: 1,
				Values: []string{
					"Default Title",
				},
			},
		},
	}

	expectedProductCSV := []shopify.ProductCSV{
		{
			Handle:                    "handle",
			Title:                     "title",
			BodyHTML:                  "bodyHTML",
			Vendor:                    "vendor",
			Type:                      "productType",
			Tags:                      "tag1,tag2",
			Published:                 "TRUE",
			Option1Name:               "Title",
			Option1Value:              "Default Title",
			VariantSKU:                "sku",
			VariantGrams:              "100",
			VariantInventoryTracker:   "shopify",
			VariantInventoryQty:       "1",
			VariantInventoryPolicy:    "deny",
			VariantFulfillmentService: "manual",
			VariantPrice:              "1000",
			VariantCompareAtPrice:     "10000",
			VariantRequiresShipping:   "TRUE",
			VariantTaxable:            "TRUE",
			VariantBarcode:            "barcode",
			GiftCard:                  "FALSE",
		},
	}

	actualProductCSV := suite.productCSVConversionService.ConvertToCSVFormat(productWithAllInformation)

	suite.Require().Equal(expectedProductCSV, actualProductCSV)
}
