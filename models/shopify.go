package models

import (
	"strconv"
	"strings"
	"time"
)

type Variant struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	Option1          string    `json:"option1"`
	Option2          string    `json:"option2"`
	Option3          string    `json:"option3"`
	Sku              string    `json:"sku"`
	RequiresShipping bool      `json:"requires_shipping"`
	Taxable          bool      `json:"taxable"`
	FeaturedImage    Image     `json:"featured_image"`
	Available        bool      `json:"available"`
	Price            string    `json:"price"`
	Grams            int       `json:"grams"`
	CompareAtPrice   string    `json:"compare_at_price"`
	Position         int       `json:"position"`
	ProductID        int64     `json:"product_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

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

type Option struct {
	Name     string   `json:"name"`
	Position int      `json:"position"`
	Values   []string `json:"values"`
}

type Product struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Handle      string    `json:"handle"`
	BodyHTML    string    `json:"body_html"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Vendor      string    `json:"vendor"`
	ProductType string    `json:"product_type"`
	Tags        []string  `json:"tags"`
	Variants    []Variant `json:"variants"`
	Images      []Image   `json:"images"`
	Options     []Option  `json:"options"`
}

// ProductsResponse holds shopify response
type ProductsResponse struct {
	Products []Product `json:"products"`
}

func (product Product) ToImportableCSV() [][]string {
	var csvRows [][]string
	for _, variant := range product.Variants {
		published := strconv.FormatBool(true)
		variantInventoryTracker := ""
		variantInventoryQuantity := "1"
		variantInventoryPolicy := "deny"
		variantFulfilmentService := "manual"
		variantBarCode := ""
		imageAltText := ""
		giftCard := strconv.FormatBool(false)
		content := []string{product.Handle, product.Title, product.BodyHTML, product.Vendor, product.ProductType, strings.Join(product.Tags, ","), published, variant.Option1, variant.Option1, variant.Option2, variant.Option2, variant.Option3, variant.Option3, variant.Sku, strconv.Itoa(variant.Grams), variantInventoryTracker, variantInventoryQuantity, variantInventoryPolicy, variantFulfilmentService, variant.Price, variant.CompareAtPrice, strconv.FormatBool(variant.RequiresShipping), strconv.FormatBool(variant.Taxable), variantBarCode, product.Images[0].Src, strconv.Itoa(product.Images[0].Position), imageAltText, giftCard}
		csvRows = append(csvRows, content)
	}
	return csvRows
}
