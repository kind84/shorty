package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const domain = "https://stick.it"

type Service struct {
	app App
}

func NewService(app App) *Service {
	return &Service{app}
}

// Bounce performs the redirect from a shortened URL to its extended version.
func (s *Service) Bounce(c *gin.Context) {
	url, err := s.app.MatchHash(c.Request.Context(), c.Params.ByName("hash"))
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

type CutRequest struct {
	URL string `json:"url"`
}

type CutResponse struct {
	ShortURL string `json:"shortUrl"`
}

// Cut returns the shortened version of the provided URL.
func (s *Service) Cut(c *gin.Context) {
	var cr CutRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&cr); err != nil {
		c.Error(err)
	}

	// check if valid url
	_, err := url.ParseRequestURI(cr.URL)
	if err != nil {
		c.Error(err)
	}

	hash, err := s.app.CutURL(c.Request.Context(), cr.URL)
	if err != nil {
		c.Error(err)
	}

	// stick domain to hash and return it
	short := fmt.Sprintf("%s/%s", domain, hash)

	c.JSON(200, CutResponse{short})
}

type BurnRequest struct {
	URL string `json:"url"`
}

// Burn removes the URL and its shortened version.
func (s *Service) Burn(c *gin.Context) {
	var br BurnRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&br); err != nil {
		c.Error(err)
	}

	if err := s.app.BurnURL(c.Request.Context(), br.URL); err != nil {
		c.Error(err)
	} else {
		c.String(200, "OK")
	}
}

type InflateRequest struct {
	ShortURL string `json:"shortUrl"`
}

type InflateResponse struct {
	URL string `json:"url"`
}

// Inflate returns the extended version of the provided short URL.
func (s *Service) Inflate(c *gin.Context) {
	var ir InflateRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&ir); err != nil {
		c.Error(err)
	}

	url, err := s.app.InflateURL(c.Request.Context(), ir.ShortURL)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, InflateResponse{url})
}

type CountRequest struct {
	URL string `json:"url"`
}

type CountResponse struct {
	Hits int `json:"hits"`
}

// Count returns the number of times a given short url has been served.
func (s *Service) Count(c *gin.Context) {
	var cr CountRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&cr); err != nil {
		c.Error(err)
	}

	n, err := s.app.CountHits(c.Request.Context(), cr.URL)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, CountResponse{n})
}
