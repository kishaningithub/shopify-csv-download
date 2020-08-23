package resources

import (
	"encoding/json"
	"fmt"
	"github.com/kishaningithub/shopify-csv-download/internal/models/shopify"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

type ShopifyResource interface {
	GetProducts(noOfRecordsPerPage int, page int) (shopify.ProductsResponse, error)
}

type shopifyResource struct {
	storeUrl   url.URL
	httpClient *http.Client
}

func NewShopifyResource(storeUrl url.URL) ShopifyResource {
	return &shopifyResource{
		storeUrl: storeUrl,
		httpClient: &http.Client{
			Timeout: time.Minute * 1,
		},
	}
}

func (resource *shopifyResource) GetProducts(noOfRecordsPerPage int, page int) (shopify.ProductsResponse, error) {
	productsJsonUrl := resource.storeUrl
	productsJsonUrl.Path = path.Join(resource.storeUrl.Path, "products.json")
	queryParams := productsJsonUrl.Query()
	queryParams.Set("limit", strconv.Itoa(noOfRecordsPerPage))
	queryParams.Set("page", strconv.Itoa(page))
	productsJsonUrl.RawQuery = queryParams.Encode()
	urlWithQueryParams := productsJsonUrl.String()
	request, err := http.NewRequest(http.MethodGet, urlWithQueryParams, nil)
	var responseForPage shopify.ProductsResponse
	if err != nil {
		return shopify.ProductsResponse{}, fmt.Errorf("unable to form products request: %w", err)
	}
	response, err := resource.httpClient.Do(request)
	if err != nil {
		return shopify.ProductsResponse{}, fmt.Errorf("unable to fetch products: %w", err)
	}
	if response.StatusCode == http.StatusNotFound {
		return shopify.ProductsResponse{}, fmt.Errorf("the url %s is not found, this could be because the site is either not built using shopify or the site has not exposed the url", urlWithQueryParams)
	}
	if response.StatusCode != http.StatusOK {
		resource.handleAPIErrorsUsingAppropriateDelays(response.StatusCode)
		return resource.GetProducts(noOfRecordsPerPage, page)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return shopify.ProductsResponse{}, fmt.Errorf("unable to read products response: %w", err)
	}
	if err = json.Unmarshal(data, &responseForPage); err != nil {
		return shopify.ProductsResponse{}, fmt.Errorf("unable to parse json: %w", err)
	}
	return responseForPage, nil
}

func (resource *shopifyResource) handleAPIErrorsUsingAppropriateDelays(httpStatusCode int) {
	if resource.is4XXResponseCode(httpStatusCode) {
		time.Sleep(2 * time.Minute) // Handle Rate Limiters
	} else if resource.is5XXResponseCode(httpStatusCode) {
		time.Sleep(20 * time.Second) // Handle internal server errors
	}
}

func (resource *shopifyResource) is4XXResponseCode(httpStatusCode int) bool {
	return httpStatusCode >= http.StatusBadRequest && httpStatusCode < http.StatusInternalServerError
}

func (resource *shopifyResource) is5XXResponseCode(httpStatusCode int) bool {
	return httpStatusCode >= http.StatusInternalServerError && httpStatusCode < 600
}
