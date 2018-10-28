package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/kishaningithub/shopify-csv-download/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	writer := csv.NewWriter(bufio.NewWriter(os.Stdout))
	defer writer.Flush()
	header := []string{"Handle", "Title", "Body (HTML)", "Vendor", "Type", "Tags", "Published", "Option1 Name", "Option1 Value", "Option2 Name", "Option2 Value", "Option3 Name", "Option3 Value", "Variant SKU", "Variant Grams", "Variant Inventory Tracker", "Variant Inventory Qty", "Variant Inventory Policy", "Variant Fulfillment Service", "Variant Price", "Variant Compare At Price", "Variant Requires Shipping", "Variant Taxable", "Variant Barcode", "Image Src", "Image Position", "Image Alt Text", "Gift Card", "Google Shopping / MPN", "Google Shopping / Age Group", "Google Shopping / Gender", "Google Shopping / Google Product Category", "SEO Title", "SEO Description", "Google Shopping / AdWords Grouping", "Google Shopping / AdWords Labels", "Google Shopping / Condition", "Google Shopping / Custom Product", "Google Shopping / Custom Label 0", "Google Shopping / Custom Label 1", "Google Shopping / Custom Label 2", "Google Shopping / Custom Label 3", "Google Shopping / Custom Label 4", "Variant Image", "Variant Weight Unit"}
	writer.Write(header)
	baseUrl := os.Args[1]
	pageNo := 1
	for {
		fullUrl := fmt.Sprintf("%s/products.json?page=%d", baseUrl, pageNo)
		response, err := http.Get(fullUrl)
		handleError("Error while fetching products", err)
		data, err := ioutil.ReadAll(response.Body)
		handleError("Error while reading product response", err)
		shopifyResponse := models.ProductsResponse{}
		err = json.Unmarshal(data, &shopifyResponse)
		if err != nil {
			os.Stderr.WriteString("Shopify is blocking requests. Waiting for a couple of minutes before the next request")
			time.Sleep(2 * time.Minute)
			continue
		}
		if len(shopifyResponse.Products) == 0 {
			return
		}
		for _, product := range shopifyResponse.Products {
			writer.WriteAll(product.ToImportableCSV())
		}
		pageNo++
	}

}

func handleError(message string, err error) {
	if err != nil {
		log.Fatal(message, " because ",	 err)
	}
}
