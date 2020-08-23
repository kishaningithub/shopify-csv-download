package shopify

import (
	"time"
)

// Variant represents a shopify variant
type Variant struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Option1          string    `json:"option1"`
	Option2          string    `json:"option2"`
	Option3          string    `json:"option3"`
	Sku              string    `json:"sku"`
	RequiresShipping bool      `json:"requires_shipping"`
	Taxable          bool      `json:"taxable"`
	Barcode          string    `json:"barcode"`
	FeaturedImage    Image     `json:"featured_image"`
	Available        bool      `json:"available"`
	Price            string    `json:"price"`
	Grams            int       `json:"grams"`
	CompareAtPrice   string    `json:"compare_at_price"`
	Position         int       `json:"position"`
	ProductID        int       `json:"product_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Image represents a shopify image
type Image struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Position   int       `json:"position"`
	UpdatedAt  time.Time `json:"updated_at"`
	ProductID  int       `json:"product_id"`
	VariantIds []int     `json:"variant_ids"`
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
	ID          int       `json:"id"`
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

type ProductCSV struct {
	Handle                              string `csv:"Handle"`
	Title                               string `csv:"Title"`
	BodyHTML                            string `csv:"Body (HTML)"`
	Vendor                              string `csv:"Vendor"`
	Type                                string `csv:"Type"`
	Tags                                string `csv:"Tags"`
	Published                           string `csv:"Published"`
	Option1Name                         string `csv:"Option1 Name"`
	Option1Value                        string `csv:"Option1 Value"`
	Option2Name                         string `csv:"Option2 Name"`
	Option2Value                        string `csv:"Option2 Value"`
	Option3Name                         string `csv:"Option3 Name"`
	Option3Value                        string `csv:"Option3 Value"`
	VariantSKU                          string `csv:"Variant SKU"`
	VariantGrams                        string `csv:"Variant Grams"`
	VariantInventoryTracker             string `csv:"Variant Inventory Tracker"`
	VariantInventoryQty                 string `csv:"Variant Inventory Qty"`
	VariantInventoryPolicy              string `csv:"Variant Inventory Policy"`
	VariantFulfillmentService           string `csv:"Variant Fulfillment Service"`
	VariantPrice                        string `csv:"Variant Price"`
	VariantCompareAtPrice               string `csv:"Variant Compare At Price"`
	VariantRequiresShipping             string `csv:"Variant Requires Shipping"`
	VariantTaxable                      string `csv:"Variant Taxable"`
	VariantBarcode                      string `csv:"Variant Barcode"`
	ImageSrc                            string `csv:"Image Src"`
	ImagePosition                       string `csv:"Image Position"`
	ImageAltText                        string `csv:"Image Alt Text"`
	GiftCard                            string `csv:"Gift Card"`
	GoogleShoppingMPN                   string `csv:"Google Shopping / MPN"`
	GoogleShoppingAgeGroup              string `csv:"Google Shopping / Age Group"`
	GoogleShoppingGender                string `csv:"Google Shopping / Gender"`
	GoogleShoppingGoogleProductCategory string `csv:"Google Shopping / Google Product Category"`
	SEOTitle                            string `csv:"SEO Title"`
	SEODescription                      string `csv:"SEO Description"`
	GoogleShoppingAdWordsGrouping       string `csv:"Google Shopping / AdWords Grouping"`
	GoogleShoppingAdWordsLabels         string `csv:"Google Shopping / AdWords Labels"`
	GoogleShoppingCondition             string `csv:"Google Shopping / Condition"`
	GoogleShoppingCustomProduct         string `csv:"Google Shopping / Custom Product"`
	GoogleShoppingCustomLabel0          string `csv:"Google Shopping / Custom Label 0"`
	GoogleShoppingCustomLabel1          string `csv:"Google Shopping / Custom Label 1"`
	GoogleShoppingCustomLabel2          string `csv:"Google Shopping / Custom Label 2"`
	GoogleShoppingCustomLabel3          string `csv:"Google Shopping / Custom Label 3"`
	GoogleShoppingCustomLabel4          string `csv:"Google Shopping / Custom Label 4"`
	VariantImage                        string `csv:"Variant Image"`
	VariantWeightUnits                  string `csv:"Variant Weight Unit"`
}
