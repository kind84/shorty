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
	err := json.NewDecoder(c.Request.Body).Decode(&cr)
	if err != nil {
		c.Error(err)
	}

	// check if valid url
	_, err = url.ParseRequestURI(cr.URL)
	if err != nil {
		c.Error(err)
	}

	hash, err := s.app.CutURL(c.Request.Context(), cr.URL)
	if err != nil {
		c.Error(err)
	}

	// stick domain to hash and return it
	short := fmt.Sprintf("%s/%s", domain, hash)

	err = json.NewEncoder(c.Copy().Writer).Encode(CutResponse{short})
	if err != nil {
		c.Error(err)
	}
}

type BurnRequest struct {
	URL string `json:"url"`
}

// Burn removes the URL and its shortened version.
func (s *Service) Burn(c *gin.Context) {
	var br BurnRequest
	err := json.NewDecoder(c.Request.Body).Decode(&br)
	if err != nil {
		c.Error(err)
	}

	err = s.app.BurnURL(c.Request.Context(), br.URL)
	if err != nil {
		c.Error(err)
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
	err := json.NewDecoder(c.Request.Body).Decode(&ir)
	if err != nil {
		c.Error(err)
	}

	url, err := s.app.InflateURL(c.Request.Context(), ir.ShortURL)
	if err != nil {
		c.Error(err)
	}

	err = json.NewEncoder(c.Copy().Writer).Encode(CutResponse{url})
	if err != nil {
		c.Error(err)
	}
}
