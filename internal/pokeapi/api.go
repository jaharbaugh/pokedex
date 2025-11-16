package pokeapi

import (
	"cache/internal/pokecache"
	"net/http"
	"io"
)

type Client struct {
	baseURL string
	http	*http.Client
	cache	*pokecache.Cache
}

func NewClient(URL string, c *pokecache.Cache) *Client {
	return &Client{
		baseURL: URL, 
		http: &http.Client{},
		cache: c,
	}
	
}

func (c *Client) get(url string) ([]byte, error) {
	if val, ok := c.cache.Get(url){
		return val, nil

	}
	
	res, err := c.http.Get(url)
	if err != nil{
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}

	c.cache.Add(url, body)
	return body, nil
}