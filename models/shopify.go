package models

import (
	"strconv"
	"strings"
	"time"
)

// Variant represents a shopify variant
type Variant struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	Option1           string    `json:"option1"`
	Option2           string    `json:"option2"`
	Option3           string    `json:"option3"`
	Sku               string    `json:"sku"`
	RequiresShipping  bool      `json:"requires_shipping"`
	Taxable           bool      `json:"taxable"`
	FeaturedImage     Image     `json:"featured_image"`
	Available         *bool     `json:"available"`
	InventoryQuantity *int      `json:"inventory_quantity"`
	Price             string    `json:"price"`
	Grams             int       `json:"grams"`
	CompareAtPrice    string    `json:"compare_at_price"`
	Position          int       `json:"position"`
	ProductID         int64     `json:"product_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// Image represents a shopify image
type Image struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Position   int       `json:"position"`
	UpdatedAt  time.Time `json:"updated_at"`
	ProductID  int64     `json:"product_id"`
	VariantIds []int64   `json:"variant_ids"`
	Src        string    `json:"src"`
	Width      int       `json:"width"`
	Height     int       `json:"height"`
}

// Option represents a shopify option
type Option struct {
	Name     string   `json:"name"`
	Position int      `json:"position"`
	Values   []string `json:"values"`
}

// Product represents a shopify product
type Product struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Handle      string      `json:"handle"`
	BodyHTML    string      `json:"body_html"`
	PublishedAt time.Time   `json:"published_at"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Vendor      string      `json:"vendor"`
	ProductType string      `json:"product_type"`
	Tags        interface{} `json:"tags"`
	Variants    []Variant   `json:"variants"`
	Images      []Image     `json:"images"`
	Options     []Option    `json:"options"`
}

// Tags is sometimes a string array and sometimes a CSV
func (product Product) GetTags() []string {
	if tags, ok := product.Tags.([]string); ok {
		return tags
	}
	if tags, ok := product.Tags.(string); ok {
		return strings.Split(tags, ", ") // yes there is an annoying space after the comma
	}
	return []string{}
}

// ProductsResponse holds shopify response
type ProductsResponse struct {
	Products []Product `json:"products"`
}

func (product Product) ToImportableCSV() [][]string {
	var csvRows [][]string
	for _, variant := range product.Variants {
		published := strconv.FormatBool(true)
		variantInventoryTracker := "shopify"
		variantInventoryQuantity := "0"
		if variant.Available != nil && *variant.Available {
			variantInventoryQuantity = "1"
		}
		if variant.InventoryQuantity != nil {
			variantInventoryQuantity = strconv.Itoa(*variant.InventoryQuantity)
		}
		variantInventoryPolicy := "deny"
		variantFulfilmentService := "manual"
		variantBarCode := ""
		imageAltText := ""
		giftCard := strconv.FormatBool(false)
		imageSrc := ""
		imagePosition := ""
		if len(product.Images) > 0 {
			imageSrc = product.Images[0].Src
			imagePosition = strconv.Itoa(product.Images[0].Position)
		}
		googleShoppingMPN := ""
		googleShoppingAgeGroup := ""
		googleShoppingGender := ""
		googleShoppingGoogleProductCategory := ""
		seoTitle := ""
		seoDescription := ""
		googleShoppingAdWordsGrouping := ""
		googleShoppingAdWordsLabels := ""
		googleShoppingCondition := ""
		googleShoppingCustomProduct := ""
		googleShoppingCustomLabel0 := ""
		googleShoppingCustomLabel1 := ""
		googleShoppingCustomLabel2 := ""
		googleShoppingCustomLabel3 := ""
		googleShoppingCustomLabel4 := ""
		variantImage := ""
		variantWeightUnit := ""

		content := []string{
			product.Handle, product.Title, product.BodyHTML, product.Vendor, product.ProductType,
			strings.Join(product.GetTags(), ","), published, variant.Option1, variant.Option1, variant.Option2,
			variant.Option2, variant.Option3, variant.Option3, variant.Sku, strconv.Itoa(variant.Grams),
			variantInventoryTracker, variantInventoryQuantity, variantInventoryPolicy, variantFulfilmentService, variant.Price,
			variant.CompareAtPrice, strconv.FormatBool(variant.RequiresShipping), strconv.FormatBool(variant.Taxable), variantBarCode,
			imageSrc, imagePosition, imageAltText, giftCard, googleShoppingMPN, googleShoppingAgeGroup, googleShoppingGender,
			googleShoppingGoogleProductCategory, seoTitle, seoDescription, googleShoppingAdWordsGrouping, googleShoppingAdWordsLabels,
			googleShoppingCondition, googleShoppingCustomProduct, googleShoppingCustomLabel0, googleShoppingCustomLabel1,
			googleShoppingCustomLabel2, googleShoppingCustomLabel3, googleShoppingCustomLabel4, variantImage, variantWeightUnit,
		}

		csvRows = append(csvRows, content)
	}
	return csvRows
}
