package shopify_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/kishaningithub/shopify-csv-download/internal/shopify"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
	"time"
)

var _ suite.SetupTestSuite = &ResourceTestSuite{}
var _ suite.TearDownTestSuite = &ResourceTestSuite{}

type ResourceTestSuite struct {
	suite.Suite
	productsJsonUrl url.URL
	resource        shopify.Resource
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}

func (suite *ResourceTestSuite) SetupTest() {
	productsJsonUrl, _ := url.Parse("https://example.com")
	suite.productsJsonUrl = *productsJsonUrl
	suite.resource = shopify.NewResource(suite.productsJsonUrl)
	httpmock.Activate()
}

func (suite *ResourceTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *ResourceTestSuite) TestGetProducts_ShouldFetchProductsAsPerGivenCriteria() {
	httpmock.RegisterResponder("GET", "https://example.com/products.json?limit=1&page=1",
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
	expectedProductsResponse := shopify.ProductsResponse{
		Products: []shopify.Product{
			{
				ID:          4409426968663,
				Title:       "Awesome product",
				Handle:      "awesome-product",
				BodyHTML:    "An amazing product",
				PublishedAt: getTime("2020-02-07T18:41:14+05:30"),
				CreatedAt:   getTime("2020-02-07T18:41:14+05:30"),
				UpdatedAt:   getTime("2020-03-07T11:21:56+05:30"),
				Vendor:      "Awesome vendor",
				ProductType: "pendants",
				Tags:        []string{"Modern Pendants"},
				Variants: []shopify.Variant{
					{
						ID:               31491876913239,
						Title:            "Default Title",
						Option1:          "Default Title",
						Sku:              "51919PA-6W LED",
						RequiresShipping: true,
						Taxable:          true,
						Available:        true,
						Price:            "6169.00",
						Grams:            5000,
						CompareAtPrice:   "0.00",
						Position:         1,
						ProductID:        4409426968663,
						CreatedAt:        getTime("2020-02-07T18:41:14+05:30"),
						UpdatedAt:        getTime("2020-03-06T13:47:58+05:30"),
					},
				},
				Images: []shopify.Image{
					{
						ID:         13653071167575,
						CreatedAt:  getTime("2020-02-08T17:58:15+05:30"),
						Position:   1,
						UpdatedAt:  getTime("2020-02-12T18:27:13+05:30"),
						ProductID:  4409426968663,
						VariantIds: []int{},
						Src:        "https://example.com/s/files/1/1204/7448/products/51919PA-6W_LED.1.jpg?v=1581512233",
						Width:      2160,
						Height:     2880,
					},
				},
				Options: []shopify.Option{
					{
						Name:     "Title",
						Position: 1,
						Values:   []string{"Default Title"},
					},
				},
			},
		},
	}

	productsResponse, err := suite.resource.GetProducts(1, 1)

	suite.Require().NoError(err)
	suite.Require().Equal(expectedProductsResponse, productsResponse)
}

func getTime(timeStr string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, timeStr)
	return parsedTime
}
