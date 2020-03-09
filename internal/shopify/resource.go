package shopify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

type Resource interface {
	GetProducts(noOfRecordsPerPage int, page int) (ProductsResponse, error)
}

type resource struct {
	storeUrl   url.URL
	httpClient *http.Client
}

func NewResource(storeUrl url.URL) Resource {
	return &resource{
		storeUrl: storeUrl,
		httpClient: &http.Client{
			Timeout: time.Minute * 1,
		},
	}
}

func (resource *resource) GetProducts(noOfRecordsPerPage int, page int) (ProductsResponse, error) {
	productsJsonUrl := resource.storeUrl
	productsJsonUrl.Path = path.Join(resource.storeUrl.Path, "products.json")
	queryParams := productsJsonUrl.Query()
	queryParams.Set("limit", strconv.Itoa(noOfRecordsPerPage))
	queryParams.Set("page", strconv.Itoa(page))
	productsJsonUrl.RawQuery = queryParams.Encode()
	urlWithQueryParams := productsJsonUrl.String()
	request, err := http.NewRequest(http.MethodGet, urlWithQueryParams, nil)
	var responseForPage ProductsResponse
	if err != nil {
		return ProductsResponse{}, fmt.Errorf("unable to form products request: %w", err)
	} else if response, err := resource.httpClient.Do(request); err != nil {
		return ProductsResponse{}, fmt.Errorf("unable to fetch products: %w", err)
	} else if response.StatusCode != http.StatusOK {
		resource.handleAPIErrorsUsingAppropriateDelays(response.StatusCode)
		return resource.GetProducts(noOfRecordsPerPage, page)
	} else if data, err := ioutil.ReadAll(response.Body); err != nil {
		return ProductsResponse{}, fmt.Errorf("unable to read products response: %w", err)
	} else if err = json.Unmarshal(data, &responseForPage); err != nil {
		return ProductsResponse{}, fmt.Errorf("unable to parse json: %w", err)
	} else {
		return responseForPage, nil
	}
}

func (resource *resource) handleAPIErrorsUsingAppropriateDelays(httpStatusCode int) {
	if httpStatusCode >= 400 && httpStatusCode < 500 {
		time.Sleep(2 * time.Minute) // Handle Rate Limiters
	} else if httpStatusCode >= 500 && httpStatusCode < 600 {
		time.Sleep(20 * time.Second) // Handle internal errors
	}
}
