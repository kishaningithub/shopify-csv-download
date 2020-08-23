package resources_test

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/kishaningithub/shopify-csv-download/internal/models/shopify"
	"github.com/kishaningithub/shopify-csv-download/internal/resources"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

var _ suite.SetupTestSuite = &ResourceTestSuite{}
var _ suite.TearDownTestSuite = &ResourceTestSuite{}

type ResourceTestSuite struct {
	suite.Suite
	productsJsonUrl url.URL
	resource        resources.ShopifyResource
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}

func (suite *ResourceTestSuite) SetupTest() {
	productsJsonUrl, _ := url.Parse("https://example.com")
	suite.productsJsonUrl = *productsJsonUrl
	suite.resource = resources.NewShopifyResource(suite.productsJsonUrl)
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
      "handle": "awesome-product"
    }
   ]
}
`))
	expectedProductsResponse := shopify.ProductsResponse{
		Products: []shopify.Product{
			{
				Handle: "awesome-product",
			},
		},
	}

	productsResponse, err := suite.resource.GetProducts(1, 1)

	suite.Require().NoError(err)
	suite.Require().Equal(expectedProductsResponse, productsResponse)
}

func (suite *ResourceTestSuite) TestGetProducts_ShouldReturnFailureWithAppropriateLogWhenUrlIsNotFound() {
	httpmock.RegisterResponder("GET", "https://example.com/products.json?limit=1&page=1",
		httpmock.NewStringResponder(404, ""))
	expectedErr := fmt.Errorf("the url https://example.com/products.json?limit=1&page=1 is not found, this could be because the site is either not built using shopify or the site has not exposed the url")

	_, actualErr := suite.resource.GetProducts(1, 1)

	suite.Require().Equal(expectedErr, actualErr)
}
