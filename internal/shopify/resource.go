package shopify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Resource interface {
	GetProducts(noOfRecordsPerPage int, page int) (ProductsResponse, error)
}

type resource struct {
	productsResourceFullUrl url.URL
	httpClient              *http.Client
}

func NewResource(productsResourceFullUrl url.URL) Resource {
	return &resource{
		productsResourceFullUrl: productsResourceFullUrl,
		httpClient: &http.Client{
			Timeout: time.Minute * 1,
		},
	}
}

func (resource *resource) GetProducts(noOfRecordsPerPage int, page int) (ProductsResponse, error) {
	productsResourceFullUrl := resource.productsResourceFullUrl
	queryParams := productsResourceFullUrl.Query()
	queryParams.Set("limit", strconv.Itoa(noOfRecordsPerPage))
	queryParams.Set("page", strconv.Itoa(page))
	productsResourceFullUrl.RawQuery = queryParams.Encode()
	request, err := http.NewRequest(http.MethodGet, productsResourceFullUrl.String(), nil)
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
