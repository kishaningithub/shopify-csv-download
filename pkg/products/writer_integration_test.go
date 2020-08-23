package products_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/kishaningithub/shopify-csv-download/pkg/products"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

var (
	_ suite.SetupTestSuite    = (*WriterTestSuite)(nil)
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

func (suite *WriterTestSuite) TestSaveAsImportableCSV_ShouldWriteCSVToGivenWriter() {
	httpmock.RegisterResponder("GET", "https://example.com/products.json?limit=250&page=1",
		httpmock.NewStringResponder(200, `
{
  "products": [
    {
      "handle": "awesome-product",
      "variants": [
        {
          "id": 31491876913239,
          "title": "Default Title",
          "option1": "Default Title"
        }
      ],
      "options": [
        {
          "name": "Title",
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
	var sb strings.Builder

	err := products.SaveAsImportableCSV("https://example.com", &sb)

	suite.Require().NoError(err)
	noOfLinesInCSV := strings.Count(sb.String(), "\n")
	suite.Require().Equal(2, noOfLinesInCSV, sb.String())
}
