package shopify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	} else if data, err := ioutil.ReadAll(response.Body); err != nil {
		log.Println("shopify is blocking requests. Waiting for a couple of minutes before the next request.")
		log.Println(err)
		time.Sleep(2 * time.Minute)
		return resource.GetProducts(noOfRecordsPerPage, page)
	} else if err = json.Unmarshal(data, &responseForPage); err != nil {
		return ProductsResponse{}, fmt.Errorf("unable to parse json: %w", err)
	} else {
		return responseForPage, nil
	}
}
