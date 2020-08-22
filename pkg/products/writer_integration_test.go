package products_test

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/kishaningithub/shopify-csv-download/pkg/products"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

var (
	_ suite.SetupTestSuite = (*WriterTestSuite)(nil)
	_ suite.TearDownTestSuite = (*WriterTestSuite)(nil)
)


type WriterTestSuite struct {
	suite.Suite
}

func TestWriterTestSuite(t *testing.T) {
	suite.Run(t, new(WriterTestSuite))
}

func (suite *WriterTestSuite) SetupTest() {
	httpmock.Activate()
}

func (suite *WriterTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *WriterTestSuite) TestWriter() {
	httpmock.RegisterResponder("GET", "https://example.com/products.json?limit=250&page=1",
		httpmock.NewStringResponder(200, `
{
  "products": [
    {
      "id": 4409426968663,
      "title": "Awesome product",
      "handle": "awesome-product",
      "body_html": "An amazing product",
      "published_at": "2020-02-07T18:41:14+05:30",
      "created_at": "2020-02-07T18:41:14+05:30",
      "updated_at": "2020-03-07T11:21:56+05:30",
      "vendor": "Awesome vendor",
      "product_type": "pendants",
      "tags": [
        "Modern Pendants"
      ],
      "variants": [
        {
          "id": 31491876913239,
          "title": "Default Title",
          "option1": "Default Title",
          "option2": null,
          "option3": null,
          "sku": "51919PA-6W LED",
          "requires_shipping": true,
          "taxable": true,
          "featured_image": null,
          "available": true,
          "price": "6169.00",
          "grams": 5000,
          "compare_at_price": "0.00",
          "position": 1,
          "product_id": 4409426968663,
          "created_at": "2020-02-07T18:41:14+05:30",
          "updated_at": "2020-03-06T13:47:58+05:30"
        }
      ],
      "images": [
        {
          "id": 13653071167575,
          "created_at": "2020-02-08T17:58:15+05:30",
          "position": 1,
          "updated_at": "2020-02-12T18:27:13+05:30",
          "product_id": 4409426968663,
          "variant_ids": [],
          "src": "https://example.com/s/files/1/1204/7448/products/51919PA-6W_LED.1.jpg?v=1581512233",
          "width": 2160,
          "height": 2880
        }
      ],
      "options": [
        {
          "name": "Title",
          "position": 1,
          "values": [
            "Default Title"
          ]
        }
      ]
    }
  ]
}
`))

	httpmock.RegisterResponder("GET", "https://example.com/products.json?limit=250&page=2",
		httpmock.NewStringResponder(200, `{"products":[]}`))
	expected := strings.TrimSpace(`
Handle,Title,Body (HTML),Vendor,Type,Tags,Published,Option1 Name,Option1 Value,Option2 Name,Option2 Value,Option3 Name,Option3 Value,Variant SKU,Variant Grams,Variant Inventory Tracker,Variant Inventory Qty,Variant Inventory Policy,Variant Fulfillment Service,Variant Price,Variant Compare At Price,Variant Requires Shipping,Variant Taxable,Variant Barcode,Image Src,Image Position,Image Alt Text,Gift Card,Google Shopping / MPN,Google Shopping / Age Group,Google Shopping / Gender,Google Shopping / Google Product Category,SEO Title,SEO Description,Google Shopping / AdWords Grouping,Google Shopping / AdWords Labels,Google Shopping / Condition,Google Shopping / Custom Product,Google Shopping / Custom Label 0,Google Shopping / Custom Label 1,Google Shopping / Custom Label 2,Google Shopping / Custom Label 3,Google Shopping / Custom Label 4,Variant Image,Variant Weight Unit
awesome-product,Awesome product,An amazing product,Awesome vendor,pendants,Modern Pendants,true,Title,Default Title,,,,,51919PA-6W LED,5000,shopify,1,deny,manual,6169.00,0.00,true,true,,https://example.com/s/files/1/1204/7448/products/51919PA-6W_LED.1.jpg?v=1581512233,1,,false,,,,,,,,,,,,,,,,,`)
	expected = fmt.Sprintln(expected)
	var sb strings.Builder

	err := products.SaveAsImportableCSV("https://example.com", &sb)

	suite.Require().NoError(err)
	suite.Require().Equal(expected, sb.String())
}
