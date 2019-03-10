package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/kishaningithub/shopify-csv-download/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var opts struct {
	FullUrl string `short:"f" long:"full-url" description:"Full URL to products.json. Eg: https://shopify-site.com/products.json"`
}

func main() {
	writer := csv.NewWriter(bufio.NewWriter(os.Stdout))
	defer writer.Flush()
	writeCSVHeader(writer)
	productsJsonURL := findProductsJsonURL()
	pageNo := 1
	log.Println("Downloading products as csv")
	for {
		fullURL := fmt.Sprintf("%s&page=%d", productsJsonURL, pageNo)
		response, err := http.Get(fullURL)
		exitWithError("Error while fetching products", err)
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println("Shopify is blocking requests. Waiting for a couple of minutes before the next request.")
			log.Println(err)
			time.Sleep(2 * time.Minute)
			continue
		}
		shopifyResponse := models.ProductsResponse{}
		err = json.Unmarshal(data, &shopifyResponse)
		exitWithError("unable to parse json", err)
		if len(shopifyResponse.Products) == 0 {
			return
		}
		for _, product := range shopifyResponse.Products {
			err := writer.WriteAll(product.ToImportableCSV())
			exitWithError("Unable to write csv", err)
		}
		pageNo++
	}
}

func findProductsJsonURL() string {
	remainingArgs, err := flags.Parse(&opts)
	exitWithError("unable to parse flags", err)
	baseURL := strings.Join(remainingArgs, "")
	productsJsonURL := fmt.Sprintf("%s/products.json?limit=250", baseURL)
	if len(opts.FullUrl) > 0 {
		productsJsonURL = opts.FullUrl
	}
	log.Println("Resolved products URL:", productsJsonURL)
	return productsJsonURL
}

func writeCSVHeader(writer *csv.Writer) {
	header := []string{"Handle", "Title", "Body (HTML)", "Vendor", "Type", "Tags", "Published", "Option1 Name", "Option1 Value", "Option2 Name", "Option2 Value", "Option3 Name", "Option3 Value", "Variant SKU", "Variant Grams", "Variant Inventory Tracker", "Variant Inventory Qty", "Variant Inventory Policy", "Variant Fulfillment Service", "Variant Price", "Variant Compare At Price", "Variant Requires Shipping", "Variant Taxable", "Variant Barcode", "Image Src", "Image Position", "Image Alt Text", "Gift Card", "Google Shopping / MPN", "Google Shopping / Age Group", "Google Shopping / Gender", "Google Shopping / Google Product Category", "SEO Title", "SEO Description", "Google Shopping / AdWords Grouping", "Google Shopping / AdWords Labels", "Google Shopping / Condition", "Google Shopping / Custom Product", "Google Shopping / Custom Label 0", "Google Shopping / Custom Label 1", "Google Shopping / Custom Label 2", "Google Shopping / Custom Label 3", "Google Shopping / Custom Label 4", "Variant Image", "Variant Weight Unit"}
	err := writer.Write(header)
	exitWithError("Unable to write csv header", err)
}

func exitWithError(message string, err error) {
	if err != nil {
		log.Fatal(message + " because " + err.Error())
	}
}
